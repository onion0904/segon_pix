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
    bucketName := "your-bucket-name" // Set your bucket name here
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
    bucketName := "your-bucket-name" // Set your bucket name here
    return repo.gcsClient.Bucket(bucketName).Object(objectName).Delete(ctx)
}

// AddPostedImage handles adding a posted image with uploading to GCS.
func (repo *Repository) AddPostedImage(ctx context.Context, file io.Reader, filename string, userID uint) error {
    url, err := repo.UploadImageToGCS(ctx, file, filename)
    if err != nil {
        return err
    }

    image := &models.PostedImage{
        URL:    url,
        UserID: userID,
    }
    return repo.db.Create(image).Error
}

// DeletePostedImage handles deleting a posted image and its corresponding file in GCS.
func (repo *Repository) DeletePostedImage(ctx context.Context, imageID uint) error {
    var image models.PostedImage
    if err := repo.db.First(&image, imageID).Error; err != nil {
        return err
    }

    if err := repo.DeleteImageFromGCS(ctx, image.URL); err != nil {
        return err
    }

    return repo.db.Delete(&image).Error
}