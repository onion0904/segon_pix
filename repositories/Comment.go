package repositories

import (
    "PixApp/models"
    "gorm.io/gorm"
)


func (repo *Repository) AddComment(model *models.Comment, imageID uint) error {
    return repo.db.Transaction(func(tx *gorm.DB) error {
        // 画像の存在確認
        var image models.PostedImage
        if err := tx.First(&image, imageID).Error; err != nil {
            return err
        }

        // コメントのPostedImageIDを設定
        model.PostedImageID = imageID

        // コメントを作成
        if err := tx.Create(model).Error; err != nil {
            return err
        }

        return nil // トランザクションをコミット
    })
}

func (repo *Repository) UpdateComment(commentID uint, newContent string, imageID uint) error {
    return repo.db.Transaction(func(tx *gorm.DB) error {
        // 更新するコメントを取得
        var comment models.Comment
        if err := tx.First(&comment, commentID).Error; err != nil {
            return err
        }

        // コメント内容を更新
        comment.Message = newContent
        if err := tx.Save(&comment).Error; err != nil {
            return err
        }

        // 画像が存在するかを確認（必要に応じて画像の存在確認を行います）
        var image models.PostedImage
        if err := tx.First(&image, imageID).Error; err != nil {
            return err
        }

        // トランザクションが自動的にコミットされる
        return nil
    })
}

func (repo *Repository) DeleteComment(commentID uint) error {
    return repo.db.Transaction(func(tx *gorm.DB) error {
        // コメントの取得
        var comment models.Comment
        if err := tx.First(&comment, commentID).Error; err != nil {
            return err
        }

        // コメントの削除（OnDelete:CASCADEで依存関係も削除される）
        if err := tx.Delete(&comment).Error; err != nil {
            return err
        }

        return nil // コミットされる
    })
}
