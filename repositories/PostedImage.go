package repositories

import (
	"PixApp/models"
    "fmt"
    "io"
	"context"
    "time"
	"os"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"strings"
	"net/url"
    "errors"
)


// UploadImageToGCS は画像を Google Cloud Storage にアップロードし、その URL を返します。
func (repo *Repository) UploadImageToGCS(ctx context.Context, file io.Reader, filename string) (string,string, error) {
    // .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	if bucketName == "" {
        return "","", fmt.Errorf("GCS bucket name is not set")
    }
    objectName := fmt.Sprintf("%v-%v", time.Now().Unix(), filename)
    object := repo.gcsClient.Bucket(bucketName).Object(objectName)
    wc := object.NewWriter(ctx)
    // ファイルの書き込み
    if _, err := io.Copy(wc, file); err != nil {
        wc.Close()
        return "","", err
    }
    if err := wc.Close(); err != nil {
        return "","", err
    }
    // オブジェクトの属性を取得してメディアリンクを取得(GCS のオブジェクトから直接 URL を取得します。)
    attrs, err := object.Attrs(ctx)
    if err != nil {
        return "","", err
    }
    return attrs.MediaLink, objectName, nil
}


// ensureHashtag は指定された名前のハッシュタグを取得、または存在しない場合は作成します。
func (repo *Repository) ensureHashtag(tx *gorm.DB, name string) (models.Hashtag, error) {
    var hashtag models.Hashtag
    err := tx.Where("name = ?", name).First(&hashtag).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // ハッシュタグが存在しない場合は作成
            hashtag = models.Hashtag{Name: name}
            if err := tx.Create(&hashtag).Error; err != nil {
                return models.Hashtag{}, fmt.Errorf("failed to create hashtag: %v", err)
            }
            return hashtag, nil
        } 
        return hashtag, nil
    }
    return hashtag, err
}



// DeleteImageFromGCS は、Google Cloud Storage から画像を削除します。
func (repo *Repository) DeleteImageFromGCS(ctx context.Context, imageURL string) error {
    // .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	bucketName := os.Getenv("GCS_BUCKET_NAME")
    if bucketName == "" {
        return fmt.Errorf("GCS bucket name is not set")
    }

    objectName, err := extractObjectNameFromURL(imageURL, bucketName)
    if err != nil {
        return err
    }

    return repo.gcsClient.Bucket(bucketName).Object(objectName).Delete(ctx)
}

// URLからオブジェクト名を抽出する関数
func extractObjectNameFromURL(imageURL, bucketName string) (string, error) {
    parsedURL, err := url.Parse(imageURL)
    if err != nil {
        return "", fmt.Errorf("invalid image URL: %v", err)
    }

    // パスからオブジェクト名を抽出
    parts := strings.SplitN(parsedURL.Path, "/", 3)
    if len(parts) < 3 || parts[1] != bucketName {
        return "", fmt.Errorf("invalid image URL path")
    }

    objectName := parts[2]
    return objectName, nil
}



//以下二つの関数を使う

// AddPostedImageは、GCSへのアップロードを伴う投稿画像の追加を処理します。
func (repo *Repository) AddPostedImage(ctx context.Context, file io.Reader, filename string, userID uint, hashtags []models.Hashtag) error {
    url, objectName, err := repo.UploadImageToGCS(ctx, file, filename)
    if err != nil {
        return err
    }
    tx := repo.db.Begin()
    if err := tx.Error; err != nil {
        return err
    }
    defer tx.Rollback()

    var user models.User
    if err = tx.First(&user, userID).Error; err != nil {
        return err
    }

    image := models.PostedImage{
        URL:        url,
        UserID:     userID,
        PostUser:   user,
        ObjectName: objectName,
    }
    if err = tx.Create(&image).Error; err != nil {
        return err
    }

    // ハッシュタグを確実に取得または作成
    for _, tag := range hashtags {
        hashtag, err := repo.ensureHashtag(tx, tag.Name)
        if err != nil {
            return err
        }
        // 画像にハッシュタグを関連付け
        if err = tx.Model(&image).Association("Hashtags").Append(&hashtag); err != nil {
            return err
        }
    }

    return tx.Commit().Error
}


// DeletePostedImageは、GCSの投稿画像と対応するファイルの削除を処理します。
func (repo *Repository) DeletePostedImage(ctx context.Context, imageID uint) error {
    return repo.db.Transaction(func(tx *gorm.DB) error {
        var image models.PostedImage
        if err := tx.First(&image, imageID).Error; err != nil {
            return err
        }

        // GCSから画像を削除
        if err := repo.DeleteImageFromGCS(ctx, image.URL); err != nil {
            return err
        }

        // 画像を削除（関連付けも自動的に処理される）
        if err := tx.Delete(&image).Error; err != nil {
            return err
        }

        return nil
    })
}
