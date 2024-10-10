package repositories

import (
	"PixApp/models"
	"errors"
	"gorm.io/gorm"
	"context"
    "fmt"
    "os"
)

// 与えられたidのユーザー情報を返す
func (repo *Repository) UserInfo(id uint) (*models.User, error) {
    var user models.User
    if err := repo.db.Preload("PostedImages").Preload("LikedImages").First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // またはカスタムエラーを返す
        }
        return nil, err
    }
    return &user, nil
}


// 与えられたハッシュタグの部分一致の画像のスライスを返す
func (repo *Repository) SearchImage(Qhashtag string) ([]models.PostedImage, error) {
    var images []models.PostedImage
    err := repo.db.
        Preload("PostUser").
        Preload("Likes").
        Preload("Comments").
        Preload("Hashtags").
        Joins("JOIN posted_image_hashtags ON posted_image_hashtags.posted_image_id = posted_images.id").
        Joins("JOIN hashtags ON posted_image_hashtags.hashtag_id = hashtags.id").
        Where("hashtags.Name LIKE ?", "%"+Qhashtag+"%").
        Find(&images).Error
    if err != nil {
        return nil, err
    }

    bucketName := os.Getenv("GCS_BUCKET_NAME")
    if bucketName == "" {
        return nil, fmt.Errorf("GCS_BUCKET_NAME is not set in environment variables")
    }

    return images, nil
}




func (repo *Repository) AddUser(model *models.User) error {
    if err := repo.db.Create(model).Error; err != nil {
        return fmt.Errorf("failed to add user to the database: %w", err)
    }
    return nil
}


func (repo *Repository) DeleteUser(userID uint) error {
    return repo.db.Transaction(func(tx *gorm.DB) error {
        // ユーザーを取得し、投稿画像をプリロード
        var user models.User
        if err := tx.Preload("PostedImages").First(&user, userID).Error; err != nil {
            return err
        }

        // ユーザーの投稿画像をGCSから削除
        for _, image := range user.PostedImages {
            // GCSから画像を削除
            if err := repo.DeletePostedImage(context.Background(), image.ID); err != nil {
                return err
            }
        }

        // ユーザーを削除
        if err := tx.Delete(&user).Error; err != nil {
            return err
        }

        return nil
    })
}

