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

// UploadImageToGCS は画像を Google Cloud Storage にアップロードし、そのURLを返す。
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

// ensureHashtag は指定された名前のハッシュタグを取得、または存在しない場合は作成する。
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