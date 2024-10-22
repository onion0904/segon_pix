package controllers

import (
	"PixApp/models"
	"PixApp/repositories"
	"net/http"
	"strconv"
    "log"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func (con *Controller) AddUser(c echo.Context) (uint,error) {
    user := models.User{}
    
    // リクエストボディのバインド
    if err := c.Bind(&user); err != nil {
        log.Printf("Failed to bind request data: %v", err)
        return 0,c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
    }

    // リポジトリの作成
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return 0,c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create repository"})
    }

    // リポジトリを使ってユーザーを追加
    if err := repo.AddUser(&user); err != nil {
        log.Printf("Failed to add user: %v", err)
        return 0,c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add user"})
    }

    return user.ID,nil
}

func (con *Controller) UserInfo(c echo.Context) error {
    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service Unavailable"})
    }
    user, err := repo.UserInfo(uintID)
    if err != nil {
        log.Printf("Failed to retrieve user info: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user information"})
    }

    return c.JSON(http.StatusOK, user)
}

func (con *Controller) UserInfoAuth(c echo.Context) error {
    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }
    email := c.QueryParam("email")
    password := c.QueryParam("password")

    if email == "" || password == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password are required"})
    }

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service Unavailable"})
    }

    user, err := repo.UserInfoAuth(email, password)
    if err != nil {
        log.Printf("Authentication failed: %v", err)
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
    }

    return c.JSON(http.StatusOK, user)
}

func (con *Controller) DeleteUser(c echo.Context) error {
    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    var existinguser models.User
    if err := con.db.First(&existinguser, uintID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            log.Printf("User not found: %d", uintID)
            return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
        }
        log.Printf("Database error: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user from database"})
    }

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service unavailable"})
    }

    if err := repo.DeleteUser(uintID); err != nil {
        log.Printf("Failed to delete user: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
    }

    log.Printf("User deleted successfully: %d", uintID)
    return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (con *Controller) UpdateUserIcon(c echo.Context) error {
    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    // アップロードされたファイルを取得
    file, err := c.FormFile("File")
    if err != nil {
        log.Printf("Failed to retrieve file: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "File is required"})
    }

    // ファイルを開く
    src, err := file.Open()
    if err != nil {
        log.Printf("Failed to open file: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open file"})
    }
    defer src.Close()

    // リポジトリを取得
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service unavailable"})
    }

    // アイコン画像をアップロードし、URLを取得
    url, _, err := repo.UploadImageToGCS(c.Request().Context(), src, file.Filename)
    if err != nil {
        log.Printf("Failed to upload image: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to upload image"})
    }

    // ユーザーのIconフィールドを更新
    err = repo.UpdateUserIcon(uintID, url)
    if err != nil {
        log.Printf("Failed to update user icon: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user icon"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "User icon updated successfully"})
}

func (con *Controller) UpdateUserHeader(c echo.Context) error {
    // クエリパラメータからユーザーIDを取得
    idStr := c.QueryParam("userID")
    if idStr == "" {
        log.Printf("User ID is missing in request")
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }

    // 文字列のIDをuintに変換
    idUint64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        log.Printf("Invalid user ID format: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID format"})
    }
    userID := uint(idUint64)

    // アップロードされたファイルを取得
    file, err := c.FormFile("File")
    if err != nil {
        log.Printf("Failed to retrieve file: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "File is required"})
    }

    // ファイルを開く
    src, err := file.Open()
    if err != nil {
        log.Printf("Failed to open file: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open file"})
    }
    defer src.Close()

    // リポジトリを取得
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service unavailable"})
    }

    // アイコン画像をアップロードし、URLを取得
    url, _, err := repo.UploadImageToGCS(c.Request().Context(), src, file.Filename)
    if err != nil {
        log.Printf("Failed to upload image: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to upload image"})
    }

    // ユーザーのHeaderフィールドを更新
    err = repo.UpdateUserHeader(userID, url)
    if err != nil {
        log.Printf("Failed to update user header: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user header"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "User header updated successfully"})
}

func (con *Controller) UpdateUserInfo(c echo.Context) error {
    // クエリパラメータからユーザーIDを取得
    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }
    err := con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    name := c.QueryParam("name")
    if name == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name is required"})
    }

    description := c.QueryParam("description")

    birthdayStr := c.QueryParam("birthday")
    birthday, err := strconv.Atoi(birthdayStr)
    if err != nil {
        log.Printf("Invalid birthday format: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Birthday must be a number"})
    }

    email := c.QueryParam("email")
    if email == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email is required"})
    }

    // リポジトリを取得
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service unavailable"})
    }

    // ユーザー情報を更新
    err = repo.UpdateUserInfo(uintID, name, description, email, birthday)
    if err != nil {
        log.Printf("Failed to update user info: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user info"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "User info updated successfully"})
}

