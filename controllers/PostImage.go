package controllers

import (
	"PixApp/models"
	"PixApp/repositories"
	"log"
	"net/http"
	"strconv"
	"strings"
    "os"
	"github.com/labstack/echo/v4"
    "fmt"
)


func (con *Controller) AddPostedImage(c echo.Context) error {
    // リクエストのコンテキストを取得
    ctx := c.Request().Context()

    // フォームデータからファイルを取得
    fileHeader, err := c.FormFile("File")
    if err != nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }

    // ファイルを開く
    file, err := fileHeader.Open()
    if err != nil {
        return c.NoContent(http.StatusInternalServerError) // 500エラー
    }
    defer file.Close()

    // ファイル名を取得
    filename := fileHeader.Filename

    // リポジトリを初期化
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable) // 503エラー
    }

    // ユーザーIDを取得し、uintに変換
    userID := c.QueryParam("ID")
    uintID64, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        return c.NoContent(500) // 400エラー
    }
    uintID := uint(uintID64)

    // マルチパートフォームデータを取得
    form, err := c.MultipartForm()
    if err != nil {
        return c.NoContent(http.StatusBadRequest) // 400エラー
    }

    // "Hashtags"フィールドからハッシュタグを取得
    hashtagNames, ok := form.Value["Hashtags"]
    if !ok {
        hashtagNames = []string{} // ハッシュタグがない場合は空スライス
    }

    // ハッシュタグ名をHashtag構造体のスライスに変換
    var hashtagsSlice []models.Hashtag
    for _, name := range hashtagNames {
        name = strings.TrimSpace(name)
        if name != "" {
            hashtagsSlice = append(hashtagsSlice, models.Hashtag{Name: name})
        }
    }

    // 画像をGCSにアップロード
    err = repo.AddPostedImage(ctx, file, filename, uintID, hashtagsSlice)
    if err != nil {
        return c.NoContent(http.StatusInternalServerError) // 500エラー
    }

    return c.NoContent(http.StatusNoContent) // 204エラー
}



// DeleteImage は指定された画像をGCSおよびデータベースから削除します
func (con *Controller) DeletePostedImage(c echo.Context) error {
    ctx := c.Request().Context()

    imageID := c.QueryParam("ID")
    log.Printf("Received delete request for image ID: %s", imageID)
    
    uintID64, err := strconv.ParseUint(imageID, 10, 64)
    if err != nil {
        log.Printf("Invalid image ID format: %v", err)
        return c.NoContent(http.StatusBadRequest)
    }
    uintID := uint(uintID64)

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service Unavailable"})
    }
    defer func() {
        if err := repo.Close(); err != nil {
            log.Printf("Error closing repository: %v", err)
        }
    }()

    if err := repo.DeletePostedImage(ctx, uintID); err != nil {
        log.Printf("Error in DeletePostedImage: %+v", err)
        
        if os.Getenv("DEBUG") == "true" {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": err.Error(),
                "detail": fmt.Sprintf("%+v", err),
            })
        }
        
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete image from storage"})
    }

    log.Printf("Successfully completed delete operation for image ID: %d", uintID)
    return c.NoContent(http.StatusNoContent)
}