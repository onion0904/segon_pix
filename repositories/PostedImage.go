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
    "errors"
    "cloud.google.com/go/storage"
)


type Image struct{
    ID       uint
    URL      string
}



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

    baseURL := "https://storage.googleapis.com/"
    url := fmt.Sprintf("%s%s/%s", baseURL, bucketName, objectName)
    return url, objectName, nil
}


// ensureHashtag は指定された名前のハッシュタグを取得、または存在しない場合は作成します。
func (repo *Repository) ensureHashtag(tx *gorm.DB, name string) (models.Hashtag, error) {
    var hashtag models.Hashtag
    err := tx.Where("Name = ?", name).First(&hashtag).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // ハッシュタグが存在しない場合は作成
            //この時、エラーが出るが気にしなくてよい
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


func (repo *Repository) DeleteImageFromGCS(w io.Writer,ctx context.Context, objectName string) error {
    bucketName := os.Getenv("GCS_BUCKET_NAME")
    if bucketName == "" {
        log.Printf("GCS bucket name is not set in environment variables")
        return fmt.Errorf("GCS bucket name is not set")
    }

    log.Printf("Attempting to delete object: %s from bucket: %s", objectName, bucketName)
    
    o := repo.gcsClient.Bucket(bucketName).Object(objectName)

    attrs, err := o.Attrs(ctx)
    if err != nil {
        return fmt.Errorf("object.Attrs: %w", err)
    }
    o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})
    
    if err := o.Delete(ctx); err != nil {
        return fmt.Errorf("Object(%q).Delete: %w", objectName, err)
    }
    

    fmt.Fprintf(w, "Blob %v deleted.\n", objectName)
    return nil
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


func (repo *Repository) DeletePostedImage(ctx context.Context, imageID uint) error {
    log.Printf("Starting deletion process for image ID: %d", imageID)
    
    return repo.db.Transaction(func(tx *gorm.DB) error {
        var image models.PostedImage
        if err := tx.First(&image, imageID).Error; err != nil {
            log.Printf("Failed to find image with ID %d: %v", imageID, err)
            return fmt.Errorf("failed to find image: %w", err)
        }
        log.Printf("Found image with ObjectName: %s", image.ObjectName)

        // URLではなくObjectNameを使用してGCSから削除
        if err := repo.DeleteImageFromGCS(os.Stdout,ctx, image.ObjectName); err != nil {
            log.Printf("Failed to delete image from GCS: %v", err)
            return fmt.Errorf("failed to delete image from GCS: %w", err)
        }

        if err := tx.Delete(&image).Error; err != nil {
            log.Printf("Failed to delete image from DB: %v", err)
            return fmt.Errorf("failed to delete image from DB: %w", err)
        }

        log.Printf("Successfully deleted image with ID: %d", imageID)
        return nil
    })
}




// 与えられたハッシュタグの部分一致の画像のスライスを返す
func (repo *Repository) SearchImage(Qhashtag string) ([]Image, error) {
    var images []models.PostedImage
    query := repo.db.
        Preload("PostUser").
        Preload("Likes").
        Preload("Comments").
        Preload("Hashtags").
        Joins("JOIN posted_image_hashtags ON posted_image_hashtags.posted_image_id = posted_images.id").
        Joins("JOIN hashtags ON posted_image_hashtags.hashtag_id = hashtags.id")

    if Qhashtag != "" {
        query = query.Where("hashtags.name LIKE ?", "%"+Qhashtag+"%")
    }

    err := query.
        Group("posted_images.id").
        Find(&images).Error

    if err != nil {
        return nil, err
    }

    var image []Image
    for _, img := range images {
        image = append(image, Image{
            ID:  img.ID,
            URL: img.URL,
        })
    }

    return image, nil
}



func (repo *Repository) ImageInfo(id uint) (*models.PostedImage, error) {
    var image models.PostedImage
    if err := repo.db.First(&image, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // またはカスタムエラーを返す
        }
        return nil, err
    }
    return &image, nil
}