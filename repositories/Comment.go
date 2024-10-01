package repositories

import (
    "PixApp/models"
)


// AddComment 新しいコメントをPostedImageに追加する
func (repo *Repository) AddComment (model *models.Comment,imageID uint) error {
 	// トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

	//コメントをDBに追加
	if err := tx.Create(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	// コメントを追加するPostedImageを取得
	var image models.PostedImage
    if err := tx.First(&image, imageID).Error; err != nil {
        tx.Rollback()
        return err
    }

	// コメントをPostedImageに追加
	if err := tx.Model(&image).Association("Comments").Append(&model); err != nil {
        tx.Rollback()
        return err
    }

	return tx.Commit().Error
}


// UpdateComment は指定されたコメントの内容を更新する
func (repo *Repository) UpdateComment(commentID uint, newContent string, imageID uint) error {
    // トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // 更新するコメントを取得
    var comment models.Comment
    if err := tx.First(&comment, commentID).Error; err != nil {
        tx.Rollback()
        return err
    }

    // DBのコメント内容を更新
    comment.Message = newContent
    if err := tx.Save(&comment).Error; err != nil {
        tx.Rollback()
        return err
    }

	// コメントを変更するPostedImageを取得
	var image models.PostedImage
    if err := tx.First(&image, imageID).Error; err != nil {
        tx.Rollback()
        return err
    }

	// ここで、特定のコメントが更新されたことをリフレッシュする操作
    for i, c := range image.Comments {
        if c.ID == comment.ID {
            image.Comments[i] = comment
            break
        }
    }

    // トランザクションをコミット
    return tx.Commit().Error
}



func (repo *Repository) DeleteComment (commentID uint, imageID uint) error {
	// トランザクションを開始
    tx := repo.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()
	
	//コメントを消す画像の取得
	var image models.PostedImage
    if err := tx.First(&image, imageID).Error; err != nil {
        tx.Rollback()
        return err
    }

	// 消すコメントの取得
	var comment models.Comment
    if err := tx.First(&comment, commentID).Error; err != nil {
        tx.Rollback()
        return err
    }

	//コメントをPostedImageから消す
	if err := tx.Model(&image).Association("Comments").Delete(&comment); err != nil {
        tx.Rollback()
        return err
    }

	// コメントをDBから消す
	if err := tx.Delete(&models.Comment{},commentID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}