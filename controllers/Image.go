package controllers

import (
	"PixApp/models"
	"PixApp/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
    
    "PixApp/util"
)


func (con *Controller) AddImage(c echo.Context) error {
    // リクエストのコンテキストを取得
    ctx := c.Request().Context()

    // フォームデータからファイルを取得
    fileHeader, err := c.FormFile("File")
    if err != nil {
        log.Printf("Error retrieving file: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ファイルの取得に失敗しました"})
    }

    // ファイルを開く
    file, err := fileHeader.Open()
    if err != nil {
        log.Printf("Error opening file: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "ファイルの読み込みに失敗しました"})
    }
    defer file.Close()

    // ファイル名を取得
    filename := fileHeader.Filename

    // リポジトリを初期化
    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Error initializing repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"message": "サービスが利用できません"})
    }

    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintID := uintID(userID)
    if uintID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "fail to convert userID to uintID"})
    }
    err = con.VerifyUserID(c, uintID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    // マルチパートフォームデータを取得
    form, err := c.MultipartForm()
    if err != nil {
        log.Printf("Error retrieving form data: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "フォームデータの取得に失敗しました"})
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
    err = repo.AddImage(ctx, file, filename, uintID, hashtagsSlice)
    if err != nil {
        log.Printf("Error uploading image to GCS: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "画像のアップロードに失敗しました"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "画像がアップロードされました"})
}

func (con *Controller) DeleteImage(c echo.Context) error {
    ctx := c.Request().Context()

    userID := c.QueryParam("userID")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
    }
    uintuserID := uintID(userID)
    if uintuserID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "fail to convert userID to uintID"})
    }
    err := con.VerifyUserID(c, uintuserID)
    if err!= nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user ID"})
    }

    imageID := c.QueryParam("imageID")
    if imageID == "" {
        log.Printf("Image ID is missing in request")
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Image ID is required"})
    }
    
    log.Printf("Received delete request for image ID: %s", imageID)
    
    uintID64, err := strconv.ParseUint(imageID, 10, 64)
    if err != nil {
        log.Printf("Invalid image ID format: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid image ID format"})
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

    if err := repo.DeleteImage(ctx, uintID); err != nil {
        log.Printf("Error in DeletePostedImage: %+v", err)
        
        if os.Getenv("DEBUG") == "true" {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": err.Error(),
                "detail": fmt.Sprintf("%+v", err),
            })
        }
        
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete image from storage"})
    }

    log.Printf("Successfully deleted image with ID: %d", uintID)
    return c.JSON(http.StatusOK, map[string]string{"message": "Image deleted successfully"})
}

func (con *Controller) GetImageInfo(c echo.Context) error {
    imageID := c.QueryParam("imageID")
    if imageID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Image ID is required"})
    }

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service Unavailable"})
    }

    uintID64, err := strconv.ParseUint(imageID, 10, 64)
    if err != nil {
        log.Printf("Invalid image ID format: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid image ID format"})
    }
    uintID := uint(uintID64)

    PostedImage, err := repo.ImageInfo(uintID)
    if err != nil {
        log.Printf("Error fetching image info: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get image information"})
    }

    return c.JSON(http.StatusOK, PostedImage)
}

func (con *Controller) GetImages(c echo.Context) error {
    params, err := util.ParseListQueryParams(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters"})
	}

    repo, err := repositories.NewRepository(con.db)
    if err != nil {
        log.Printf("Failed to create repository: %v", err)
        return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "サービスが利用できません"})
    }
    
    var PostedImage []models.Image

    if params.Search && params.Like{
        PostedImage, err = repo.GetLikeSearchImage(params.Hashtag,params.CurrentID,params.LikeNum)
        if err != nil {
            log.Printf("Error LikeSearching for image by hashtag: %v", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "検索した人気の画像の取得に失敗しました"})
        }
    } else if params.Search && !params.Like{
        PostedImage, err = repo.GetSearchImage(params.Hashtag,params.CurrentID)
        if err != nil {
            log.Printf("Error searching for image by hashtag: %v", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "検索した画像の取得に失敗しました"})
        }
    } else if !params.Search && params.Like{
        PostedImage, err = repo.GetLikeImages(params.CurrentID,params.LikeNum)
        if err != nil {
            log.Printf("Failed to get liked images: %v", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "人気の画像の取得に失敗しました"})
        }
    } else {
        PostedImage, err = repo.GetRecentImages(params.CurrentID)
        if err != nil {
            log.Printf("Failed to get recent images: %v", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "最近の画像の取得に失敗しました"})
        }
    }

    return c.JSON(http.StatusOK, PostedImage)
}