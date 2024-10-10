package repositories

import (
	"PixApp/models"
	"errors"
	"gorm.io/gorm"
	"context"
    "fmt"
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

func (repo *Repository) UpdateUserIcon(userID uint, iconURL string) error {
    var user models.User
    if err := repo.db.First(&user, userID).Error; err != nil {
        return err
    }

    user.Icon = iconURL

    if err := repo.db.Save(&user).Error; err != nil {
        return err
    }

    return nil
}

func (repo *Repository) UpdateUserHeader(userID uint, iconURL string) error {
    var user models.User
    if err := repo.db.First(&user, userID).Error; err != nil {
        return err
    }

    user.HeaderImage = iconURL

    if err := repo.db.Save(&user).Error; err != nil {
        return err
    }

    return nil
}