package repositories

import (
    "PixApp/models"
)

func (repo *Repository) AddFollow(followingID uint,followedID uint) error {
    // トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // ユーザーを取得
    var followinguser models.User
    if err := tx.First(&followinguser, followingID).Error; err != nil {
        tx.Rollback()
        return err
    }

	// ユーザーを取得
    var followeduser models.User
    if err := tx.First(&followeduser, followedID).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Model(&followinguser).Association("Follows").Append(&followeduser); err != nil {
        tx.Rollback()
        return err
    }
    // トランザクションをコミット
    return tx.Commit().Error
}

func (repo *Repository) RemoveFollow(followingID uint,followedID uint) error {
    // トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // ユーザーを取得
    var followinguser models.User
    if err := tx.First(&followinguser, followingID).Error; err != nil {
        tx.Rollback()
        return err
    }

	// ユーザーを取得
    var followeduser models.User
    if err := tx.First(&followeduser, followedID).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Model(&followinguser).Association("Follows").Delete(&followeduser); err != nil {
        tx.Rollback()
        return err
    }

    // トランザクションをコミット
    return tx.Commit().Error
}
