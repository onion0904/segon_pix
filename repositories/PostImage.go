package repositories

import (
	"PixApp/models"
    "fmt"
    "io"
	"context"
    "time"
)


// UploadImageToGCS は画像を Google Cloud Storage にアップロードし、その URL を返します。
func (repo *Repository) UploadImageToGCS(ctx context.Context, file io.Reader, filename string) (string, error) {
    bucketName := "GCS_BUCKET_NAME"
	if bucketName == "" {
        return "", fmt.Errorf("GCS bucket name is not set")
    }
    objectName := fmt.Sprintf("%v-%v", time.Now().Unix(), filename) // Unique object name
    wc := repo.gcsClient.Bucket(bucketName).Object(objectName).NewWriter(ctx)
    if _, err := io.Copy(wc, file); err != nil {
        return "", err
    }
    if err := wc.Close(); err != nil {
        return "", err
    }

    return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName), nil
}

// DeleteImageFromGCS は、Google Cloud Storage から画像を削除します。
func (repo *Repository) DeleteImageFromGCS(ctx context.Context, objectName string) error {
    bucketName := "GCS_BUCKET_NAME"
	if bucketName == "" {
        return fmt.Errorf("GCS bucket name is not set")
    }
    return repo.gcsClient.Bucket(bucketName).Object(objectName).Delete(ctx)
}



//以下二つの関数を使う

// AddPostedImageは、GCSへのアップロードを伴う投稿画像の追加を処理します。
func (repo *Repository) AddPostedImage(ctx context.Context, file io.Reader, filename string, userID uint,hashtags []models.Hashtag) error {
    url, err := repo.UploadImageToGCS(ctx, file, filename)
    if err != nil {
        return err
    }

	tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

	var user models.User
    if err := tx.First(&user, userID).Error; err != nil {
        tx.Rollback()
        return err
    }

    image := &models.PostedImage{
        URL:    url,
        UserID: userID,
		PostUser: user,
		Hashtags: hashtags,
    }
    if err := repo.db.Create(image).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&user).Association("PostedImages").Append(&image); err != nil {
        tx.Rollback()
        return err
    }

	return tx.Commit().Error
}

// DeletePostedImageは、GCSの投稿画像と対応するファイルの削除を処理します。
func (repo *Repository) DeletePostedImage(ctx context.Context, imageID uint) error {
    var image models.PostedImage
    if err := repo.db.First(&image, imageID).Error; err != nil {
        return err
    }

    if err := repo.DeleteImageFromGCS(ctx, image.URL); err != nil {
        return err
    }

	tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

	var user models.User
    if err := tx.First(&user, image.UserID).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Delete(&image).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&user).Association("PostedImages").Delete(&image); err != nil {
        tx.Rollback()
        return err
    }

	return tx.Commit().Error
}