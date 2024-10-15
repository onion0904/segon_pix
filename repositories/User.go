package repositories

import (
	"PixApp/models"
	"errors"
	"gorm.io/gorm"
	"context"
    "fmt"
    "golang.org/x/crypto/bcrypt"
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

    user.Email = ""
    user.Password = ""

    return &user, nil
}


// 認証付きでユーザー情報を返す
func (repo *Repository) UserInfoAuth(email, password string) (*models.User, error) {
    var user models.User
    // Emailでユーザーを検索
    if err := repo.db.Preload("PostedImages").Preload("LikedImages").
        Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }

    // パスワードの照合
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid password")
    }

    return &user, nil
}


func (repo *Repository) AddUser(user *models.User) error {
    // パスワードをハッシュ化する
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }
    user.Password = string(hashedPassword)

    // ユーザーをデータベースに追加
    if err := repo.db.Create(user).Error; err != nil {
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


func (repo *Repository) UpdateUserInfo(userID uint, name string, description string, email string) error {
    var user models.User
    if err := repo.db.First(&user, userID).Error; err != nil {
        return err
    }

    user.Name = name
    user.Description = description
    user.Email = email

    if err := repo.db.Save(&user).Error; err != nil {
        return err
    }

    return nil
}