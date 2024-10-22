package repositories

import (
	"PixApp/models"
    "fmt"
    "io"
	"context"
	"os"
	"log"
	"gorm.io/gorm"
    "errors"
    "gorm.io/gorm/clause"
)

//以下二つの関数を使う
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

func (repo *Repository) SearchImage(Qhashtag string) ([]models.Image, error) {
    var images []models.PostedImage
    query := repo.db.
        Preload(clause.Associations).
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

    var image []models.Image
    for _, img := range images {
        image = append(image, models.Image{
            ID:  img.ID,
            URL: img.URL,
        })
    }

    return image, nil
}

func (repo *Repository) ImageInfo(id uint) (*models.PostedImage, error) {
    var image models.PostedImage
    if err := repo.db.
    Preload(clause.Associations).
    First(&image, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // またはカスタムエラーを返す
        }
        return nil, err
    }
    return &image, nil
}

func (repo *Repository) GetRecentImages() ([]models.Image, error) {
    var postedImages []models.PostedImage
    err := repo.db.
        Preload(clause.Associations).
        Order("posted_images.created_at DESC"). // 作成日時の降順で並び替え
        Find(&postedImages).Error

    if err != nil {
        return nil, err
    }

    var images []models.Image
    for _, img := range postedImages {
        images = append(images, models.Image{
            ID:  img.ID,
            URL: img.URL,
        })
    }

    return images, nil
}

func (repo *Repository) GetLikeImages() ([]models.Image, error) {
    var postedImages []models.PostedImage
    err := repo.db.
        Preload(clause.Associations).
        Select("posted_images.*, COUNT(posted_image_likes.user_id) as likes_count").
        Joins("LEFT JOIN posted_image_likes ON posted_image_likes.posted_image_id = posted_images.id").
        Group("posted_images.id").
        Order("likes_count DESC"). // いいね数の降順で並び替え
        Find(&postedImages).Error

    if err != nil {
        return nil, err
    }

    var images []models.Image
    for _, img := range postedImages {
        images = append(images, models.Image{
            ID:  img.ID,
            URL: img.URL,
        })
    }
    return images, nil
}

