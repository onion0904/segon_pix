package controllers

import (
	"PixApp/models"
	"PixApp/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func (con *Controller) AddUser(c echo.Context) error {
    user := models.User{}
    
    // リクエストボディのバインド
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request data")
    }
    if user.Followers == nil {
        user.Followers = make([]models.User, 0)
    }
    if user.Follows == nil {
        user.Follows = make([]models.User, 0)
    }

    // リポジトリの作成
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create repository")
    }

    // リポジトリを使ってユーザーを追加
    if err := repo.AddUser(&user); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to add user")
    }

    return c.JSON(http.StatusOK, nil) // 正常終了
}


func (con *Controller) UserInfo(c echo.Context) error {
	userID := c.QueryParam("ID")
	uintID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintID := uint(uintID64)
	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	User, err := repo.UserInfo(uintID)
	if err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.JSON(http.StatusOK, User)
}


func (con *Controller) UserInfoAuth(c echo.Context) error {
    email := c.QueryParam("email")
    password := c.QueryParam("password")

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.JSON(http.StatusServiceUnavailable,err) // 503エラー
    }

    user ,err := repo.UserInfoAuth(email,password)
	if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "メールアドレスまたはパスワードが間違っています"})
    }

    return c.JSON(http.StatusOK, user)
}


func (con *Controller) DeleteUser(c echo.Context) error {
	userID := c.QueryParam("ID")
	uintID64, err := strconv.ParseUint(userID, 10, 64)
	if err!= nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }
	uintID := uint(uintID64)

	var existinguser models.User
	if err := con.db.First(&existinguser, uintID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound) // 404エラー
		}
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}

	repo ,err := repositories.NewRepository(con.db)
	if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }
	if err := repo.DeleteUser(uintID); err != nil {
		return c.NoContent(http.StatusServiceUnavailable) // 503エラー
	}
	return c.NoContent(http.StatusOK)
}


func (con *Controller) UpdateUserIcon(c echo.Context) error {
    // クエリパラメータからユーザーIDを取得
    idStr := c.QueryParam("ID")
    if idStr == "" {
        return c.String(http.StatusBadRequest, "IDクエリパラメータが必要です")
    }

    // 文字列のIDをuintに変換
    idUint64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return c.String(http.StatusBadRequest, "無効なIDの形式です")
    }
    userID := uint(idUint64)

    // アップロードされたファイルを取得
    file, err := c.FormFile("File")
    if err != nil {
        return c.String(http.StatusBadRequest, "ファイルが必要です")
    }

    // ファイルを開く
    src, err := file.Open()
    if err != nil {
        return c.String(http.StatusInternalServerError, "ファイルを開けませんでした")
    }
    defer src.Close()

    // リポジトリを取得
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

    // アイコン画像をアップロードし、URLを取得
    url, _, err := repo.UploadImageToGCS(c.Request().Context(), src, file.Filename)
    if err != nil {
        return c.String(http.StatusInternalServerError, "画像のアップロードに失敗しました")
    }

    // ユーザーのIconフィールドを更新
    err = repo.UpdateUserIcon(userID, url)
    if err != nil {
        return c.String(http.StatusInternalServerError, "ユーザーのアイコン更新に失敗しました")
    }

    return c.NoContent(http.StatusOK)
}


func (con *Controller) UpdateUserHeader(c echo.Context) error {
    // クエリパラメータからユーザーIDを取得
    idStr := c.QueryParam("ID")
    if idStr == "" {
        return c.String(http.StatusBadRequest, "IDクエリパラメータが必要です")
    }

    // 文字列のIDをuintに変換
    idUint64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return c.String(http.StatusBadRequest, "無効なIDの形式です")
    }
    userID := uint(idUint64)

    // アップロードされたファイルを取得
    file, err := c.FormFile("File")
    if err != nil {
        return c.String(http.StatusBadRequest, "ファイルが必要です")
    }

    // ファイルを開く
    src, err := file.Open()
    if err != nil {
        return c.String(http.StatusInternalServerError, "ファイルを開けませんでした")
    }
    defer src.Close()

    // リポジトリを取得
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

    // アイコン画像をアップロードし、URLを取得
    url, _, err := repo.UploadImageToGCS(c.Request().Context(), src, file.Filename)
    if err != nil {
        return c.String(http.StatusInternalServerError, "画像のアップロードに失敗しました")
    }

    // ユーザーのIconフィールドを更新
    err = repo.UpdateUserHeader(userID, url)
    if err != nil {
        return c.String(http.StatusInternalServerError, "ユーザーのアイコン更新に失敗しました")
    }

    return c.NoContent(http.StatusOK)
}



func (con *Controller) UpdateUserInfo(c echo.Context) error {
    // クエリパラメータからユーザーIDを取得
    idStr := c.QueryParam("userID")
    if idStr == "" {
        return c.String(http.StatusBadRequest, "IDクエリパラメータが必要です")
    }

    // 文字列のIDをuintに変換
    idUint64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return c.String(http.StatusBadRequest, "無効なIDの形式です")
    }
    userID := uint(idUint64)

	name := c.QueryParam("name")
    if name == "" {
        return c.String(http.StatusBadRequest,"nothing name")
    }

	description := c.QueryParam("description")

    birthdayStr := c.QueryParam("birthday")
    birthday, err := strconv.Atoi(birthdayStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "誕生日は数値で指定してください"})
    }


	email := c.QueryParam("email")
    if email == "" {
        return c.String(http.StatusBadRequest, "nothing email")
    }

    // リポジトリを取得
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

    // ユーザーのIconフィールドを更新
    err = repo.UpdateUserInfo(userID, name,description,email,birthday)
    if err != nil {
        return c.String(http.StatusInternalServerError, "ユーザーのアイコン更新に失敗しました")
    }

    return c.NoContent(http.StatusOK)
}
