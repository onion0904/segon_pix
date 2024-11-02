package controllers

import(
    "PixApp/mail"
	"PixApp/repositories"
    "PixApp/models"
	"crypto/rand"
    "encoding/base64"
    "fmt"
    "net/http"
    "time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
    "log"
)


func (con *Controller) Signup(c echo.Context) error {
    email := c.QueryParam("email")
    if email == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "メールアドレスが必要です"})
    }

    // ランダムな認証コードを生成
    code, err := generateVerificationCode()
    if err != nil {
        log.Printf("Error generating verification code: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "認証コードの生成に失敗しました"})
    }

    // 認証コードを保存
    con.auth.CodeMutex.Lock()
    con.auth.VerificationCodes[email] = code
    con.auth.CodeMutex.Unlock()

    // 認証コードをメールで送信（デモではコンソールに出力）
    fmt.Printf("認証コードを %s に送信しました: %s\n", email, code)
    mail.SendEmail(email, code)
    return c.JSON(http.StatusOK, map[string]string{"message": "認証コードをメールに送信しました"})
}

func generateVerificationCode() (string, error) {
    b := make([]byte, 6)
    if _, err := rand.Read(b); err != nil {
        log.Printf("Error generating random bytes: %v", err)
        return "", err
    }
    return base64.StdEncoding.EncodeToString(b), nil
}

func (con *Controller) VerifyAddUser(c echo.Context) error {
    secret := os.Getenv("JWT_SECRET_KEY")
    if secret == "" {
        log.Printf("JWT secret key is not set")
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "サーバー設定に問題があります"})
    }
    jwtSecret := []byte(secret)

    jsonData := models.User{}

    // リクエストボディのバインド
    if err := c.Bind(&jsonData); err != nil {
        log.Printf("Failed to bind request data: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
    }
    email := jsonData.Email
    password := jsonData.Password
    code := c.QueryParam("code")

    if email == "" || password == "" || code == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "全ての項目を入力してください"})
    }

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error creating repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    ok, err := repo.ExistUser(email, password)
    if err != nil {
        log.Printf("Error checking user existence: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "ユーザー確認中にエラーが発生しました"})
    }

    if ok!=0 {
        return c.JSON(http.StatusConflict, map[string]string{"message": "すでにユーザーが登録されています"})
    }

    fmt.Printf("Email is %s \n", email)

    // 認証コードの検証
    con.auth.CodeMutex.Lock()
    expectedCode, exists := con.auth.VerificationCodes[email]
    con.auth.CodeMutex.Unlock()

    if !exists {
        log.Printf("Verification code for %s not found", email)
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "認証コードが見つかりません"})
    }

    if expectedCode != code {
        log.Printf("Invalid verification code for %s", email)
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "認証コードが正しくありません"})
    }

    // 認証コードを削除
    con.auth.CodeMutex.Lock()
    delete(con.auth.VerificationCodes, email)
    con.auth.CodeMutex.Unlock()

    userID,err := con.AddUser(c)
    if err != nil{
        log.Printf("Error adding user: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "ユーザーの追加に失敗しました"})
    }

    claims := &models.MyCustomClaims{
        Email: email,
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   email,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 有効期限
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        log.Printf("Error signing JWT token: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "トークンの作成に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func (con *Controller) Login(c echo.Context) error {
    secret := os.Getenv("JWT_SECRET_KEY")
    if secret == "" {
        log.Printf("JWT secret key is not set")
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "サーバー設定に問題があります"})
    }
    jwtSecret := []byte(secret)

    email := c.QueryParam("email")
    password := c.QueryParam("password")
    if email == "" || password == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "メールアドレスとパスワードが必要です"})
    }

    // リポジトリを初期化
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error creating repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    // ユーザーの存在を確認
    userID, err := repo.ExistUser(email, password)
    if err != nil {
        log.Printf("Error checking user existence: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "ユーザー確認中にエラーが発生しました"})
    }
    if userID==0 {
        return c.JSON(http.StatusConflict, map[string]string{"message": "ユーザーが登録されていません"})
    }

    // JWTのクレーム作成
    claims := &models.MyCustomClaims{
        Email: email,
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   email,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 有効期限24時間
        },
    }

    // JWTトークンの作成
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        log.Printf("Error signing JWT token: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "トークンの作成に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func (con *Controller) Restricted(c echo.Context) error {
    userToken, ok := c.Get("user").(*jwt.Token)
    if !ok {
        log.Printf("Failed to get JWT token from context")
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "認証が必要です"})
    }

    claims, ok := userToken.Claims.(*models.MyCustomClaims)
    if !ok {
        log.Printf("Invalid claims in JWT token")
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "トークンのクレームが無効です"})
    }

    email := claims.Email
    if email == "" {
        log.Printf("No email found in JWT claims")
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "トークンにメールアドレスが含まれていません"})
    }

    return c.String(http.StatusOK, "Welcome "+email+"!")
}

func (con *Controller) VerifyUserID(c echo.Context,userID uint) error {
    userToken, ok := c.Get("user").(*jwt.Token)
    if !ok {
        log.Printf("Failed to get JWT token from context")
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "認証が必要です"})
    }

    claims, ok := userToken.Claims.(*models.MyCustomClaims)
    if !ok {
        log.Printf("Invalid claims in JWT token")
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "トークンのクレームが無効です"})
    }

    VerifyUserID := claims.UserID
    if VerifyUserID == 0 {
        log.Printf("No userID found in JWT claims")
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "トークンにuserIDが含まれていません"})
    }

    //userIDとJWTのuserIDが一致するかの確認
    if userID != VerifyUserID {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
    }
    return nil
}
