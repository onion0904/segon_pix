package repositories

import (
    "PixApp/models"
)

// AddLike はユーザーが画像にいいねをする処理
func (repo *Repository) AddLike(userID uint, imageID uint) error {
    // トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // いいねする画像を取得
    var image models.PostedImage
    if err := tx.First(&image, imageID).Error; err != nil {
        tx.Rollback()
        return err
    }

    // ユーザーを取得
    var user models.User
    if err := tx.First(&user, userID).Error; err != nil {
        tx.Rollback()
        return err
    }

    // 画像のLikesにユーザーを追加
    if err := tx.Model(&image).Association("Likes").Append(&user); err != nil {
        tx.Rollback()
        return err
    }

    // ユーザーのLikedImagesに画像を追加
    if err := tx.Model(&user).Association("LikedImages").Append(&image); err != nil {
        tx.Rollback()
        return err
    }

    // トランザクションをコミット
    return tx.Commit().Error
}

// RemoveLike はユーザーが画像のいいねを取り消す処理
func (repo *Repository) RemoveLike(userID uint, imageID uint) error {
    // トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // いいねを取り消す画像を取得
    var image models.PostedImage
    if err := tx.First(&image, imageID).Error; err != nil {
        tx.Rollback()
        return err
    }

    // ユーザーを取得
    var user models.User
    if err := tx.First(&user, userID).Error; err != nil {
        tx.Rollback()
        return err
    }

    // 画像のLikesからユーザーを削除
    if err := tx.Model(&image).Association("Likes").Delete(&user); err != nil {
        tx.Rollback()
        return err
    }

    // ユーザーのLikedImagesから画像を削除
    if err := tx.Model(&user).Association("LikedImages").Delete(&image); err != nil {
        tx.Rollback()
        return err
    }

    // トランザクションをコミット
    return tx.Commit().Error
}
