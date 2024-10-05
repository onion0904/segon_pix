package repositories

import (
	"PixApp/models"
	"errors"
	"gorm.io/gorm"
	"context"
    "strings"
)

// 与えられたidのユーザー情報を返す
func (repo *Repository) UserInfo(id uint) (*models.User, error) {
    var user models.User
    if err := repo.db.First(&user, id).Error; err != nil {
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
        Where("hashtags.name LIKE ?", "%"+Qhashtag+"%").
        Find(&images).Error
    if err != nil {
        return nil, err
    }

    // URLから "/download/" を取り除く
    for i, image := range images {
        images[i].URL = strings.Replace(image.URL, "/download/", "/", 1)
    }

    return images, nil
}




func (repo *Repository) AddUser(model *models.User) error {
    return repo.db.Create(model).Error
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

