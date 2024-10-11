package controllers

import(
	"PixApp/repositories"
	"crypto/rand"
    "encoding/base64"
    "fmt"
    "net/http"
    "sync"
    "time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
)


// 認証コードを保存するためのマップ（メールアドレスとコードの対応）
var verificationCodes = make(map[string]string)
var codeMutex sync.Mutex


func (con *Controller) Login(c echo.Context) error {

    email := c.FormValue("email")
    password := c.FormValue("password")

	// リポジトリを初期化
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

    err = repo.ExistUser(email,password)
	if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "メールアドレスまたはパスワードが間違っています"})
    }

    // ランダムな認証コードを生成
    code, err := generateVerificationCode()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "認証コードの生成に失敗しました"})
    }

    // 認証コードを保存
    codeMutex.Lock()
    verificationCodes[email] = code
    codeMutex.Unlock()

    // 認証コードをメールで送信（デモではコンソールに出力）
    fmt.Printf("認証コードを %s に送信しました: %s\n", email, code)
    // 実際のアプリケーションでは、ここでメール送信処理を行います

    return c.JSON(http.StatusOK, map[string]string{"message": "認証コードをメールに送信しました"})
}




func (con *Controller) Verify(c echo.Context) error {

	// // JWTの署名に使用するシークレットキー（実際のアプリケーションでは安全な場所に保管）
	secret := os.Getenv("JWT_SECRET_KEY")
	jwtSecret := []byte(secret)


    email := c.FormValue("email")
    code := c.FormValue("code")

    // 認証コードの検証
    codeMutex.Lock()
    expectedCode, ok := verificationCodes[email]
    codeMutex.Unlock()

    if !ok || expectedCode != code {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "認証コードが正しくありません"})
    }

    // 認証コードを削除（再利用を防ぐため）
    codeMutex.Lock()
    delete(verificationCodes, email)
    codeMutex.Unlock()

    // JWTトークンの作成
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = email
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24時間有効

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "トークンの作成に失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}


func generateVerificationCode() (string, error) {
    b := make([]byte, 6)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(b), nil
}



func (con *Controller) Restricted(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwt.RegisteredClaims)
    email := claims.Subject

    return c.String(http.StatusOK, "ようこそ "+email+" さん！")
}