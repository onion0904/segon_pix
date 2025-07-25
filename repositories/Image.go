package repositories

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"os"
	"time"

	"PixApp/models"
	"PixApp/util"
)

func (repo *Repository) AddImage(ctx context.Context, file io.Reader, filename string, userID uint, hashtags []models.Hashtag) error {
	url, objectName, err := repo.UploadImageToGCS(ctx, file, filename)
	if err != nil {
		return err
	}

	// db.Transactionメソッドを使用
	// 渡された関数内でエラーが返されれば自動でロールバック、nilが返れば自動でコミットされる
	txErr := repo.db.Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		image := models.PostedImage{
			URL:        url,
			UserID:     userID,
			PostUser:   user,
			ObjectName: objectName,
		}
		if err := tx.Create(&image).Error; err != nil {
			return err
		}

		// ハッシュタグを確実に取得または作成
		for _, tag := range hashtags {
			hashtag, err := repo.ensureHashtag(tx, tag.Name)
			if err != nil {
				return err // エラーを返せばトランザクション全体がロールバックされる
			}

			// 画像にハッシュタグを関連付け
			if err := tx.Model(&image).Association("Hashtags").Append(&hashtag); err != nil {
				return err
			}
		}

		// すべて成功した場合、nilを返すことでコミットされる
		return nil
	})

	// トランザクションが失敗（Rollback）した場合の処理
	if txErr != nil {
		log.Printf("DBトランザクションが失敗しました。GCSオブジェクトのクリーンアップを開始します。ObjectName: %s", objectName)

		// GCS画像削除のリトライ処理
		const maxRetries = 5
		var gcsDeleteErr error
		bucketName := os.Getenv("GCS_BUCKET_NAME")
		for attempt := 1; attempt <= maxRetries; attempt++ {
			gcsDeleteErr = repo.DeleteImageFromGCS(os.Stdout, ctx, objectName, bucketName)
			if gcsDeleteErr == nil {
				// 削除に成功
				log.Printf("GCSオブジェクトのクリーンアップに成功しました。ObjectName: %s", objectName)
				// 成功したので、元のDBエラーを返して終了
				return txErr
			}
			// 削除に失敗
			log.Printf("GCSオブジェクトのクリーンアップに失敗しました (試行 %d/%d)。エラー: %v", attempt, maxRetries, gcsDeleteErr)
			time.Sleep(2 * time.Second) // 2秒待ってからリトライ
		}

		// リトライがすべて失敗した場合
		if gcsDeleteErr != nil {
			// エラーメッセージを作成
			errorMessage := fmt.Sprintf("[画像の追加機能] GCSオブジェクト '%s' の自動クリーンアップに5回失敗しました。手動での対応が必要です。最終エラー: %v", objectName, gcsDeleteErr)

			// ログに致命的なエラーとして記録
			log.Println(errorMessage)

			// DBに失敗記録を残すための関数呼び出し
			err = repo.logFailToDB(objectName, gcsDeleteErr.Error())
			if err != nil {
				errorMessage = fmt.Sprintf("%s DBにログを残せませんでした: %v", errorMessage, err)
			}

			//slackに送信
			util.SlackNoticeTransaction(errorMessage)
		}

		// 呼び出し元には、そもそもの原因であるDBエラーを返す
		return txErr
	}

	return nil
}

func (repo *Repository) DeleteImage(ctx context.Context, imageID uint) error {
	log.Printf("Starting deletion process for image ID: %d", imageID)
	var image models.PostedImage
	var objectName string

	if err := repo.db.First(&image, imageID).Error; err != nil {
		log.Printf("Failed to find image with ID %d: %v", imageID, err)
		return fmt.Errorf("failed to find image: %w", err)
	}
	objectName = image.ObjectName
	log.Printf("Found image with ObjectName: %s", objectName)

	//画像を一時的な場所に移動
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	subBucketName := os.Getenv("GCS_SUB_BUCKET_NAME")
	subObjectName, err := repo.MoveImageToTemp(bucketName, subBucketName, ctx, objectName)
	if err != nil {
		return fmt.Errorf("failed to move image bucket: %w", err)
	}

	//画像をDBから削除
	deleteImageErr := repo.db.Unscoped().Delete(&image).Error
	if deleteImageErr != nil {
		log.Printf("DBトランザクションが失敗しました。GCSオブジェクトを元の場所に戻す作業をを開始します。ObjectName: %s", objectName)

		bucketName := os.Getenv("GCS_BUCKET_NAME")
		subBucketName := os.Getenv("GCS_SUB_BUCKET_NAME")
		//MoveImageToTempにリトライ、報告機能があるのでそれは省略
		_, gcsTempErr := repo.MoveTempToImage(subBucketName, bucketName, ctx, objectName)
		if gcsTempErr == nil {
			log.Printf("GCSオブジェクトの移動に成功しました。ObjectName: %s", objectName)
			// 成功したので、元のDBエラーを返して終了
			return deleteImageErr
		}
	}

	// 画像ファイルを一時的な場所から削除
	if err := repo.DeleteImageFromGCS(os.Stdout, ctx, subObjectName, subBucketName); err != nil {
		log.Printf("画像ファイルを一時的な場所から削除が失敗しました。GCSオブジェクトのクリーンアップを開始します。ObjectName: %s", subObjectName)

		// GCS画像削除のリトライ処理
		const maxRetries = 5
		var gcsDeleteErr error
		for attempt := 1; attempt <= maxRetries; attempt++ {
			gcsDeleteErr = repo.DeleteImageFromGCS(os.Stdout, ctx, subObjectName, subBucketName)
			if gcsDeleteErr == nil {
				// 削除に成功
				log.Printf("GCSオブジェクトのクリーンアップに成功しました。ObjectName: %s", subObjectName)
				// 成功したので、元のDBエラーを返して終了
				return err
			}
			// 削除に失敗
			log.Printf("GCSオブジェクトのクリーンアップに失敗しました (試行 %d/%d)。エラー: %v", attempt, maxRetries, gcsDeleteErr)
			time.Sleep(2 * time.Second) // 2秒待ってからリトライ
		}

		// リトライがすべて失敗した場合
		if gcsDeleteErr != nil {
			// エラーメッセージを作成
			errorMessage := fmt.Sprintf("[画像ファイル一時的な場所から削除] GCSオブジェクト '%s' の自動クリーンアップに5回失敗しました。手動での対応が必要です。最終エラー: %v", subObjectName, gcsDeleteErr)

			// ログに致命的なエラーとして記録
			log.Println(errorMessage)

			// DBに失敗記録を残すための関数呼び出し
			err = repo.logFailToDB(subObjectName, gcsDeleteErr.Error())
			if err != nil {
				errorMessage = fmt.Sprintf("%s DBにログを残せませんでした: %v", errorMessage, err)
			}

			//slackに送信
			util.SlackNoticeTransaction(errorMessage)
		}
		return fmt.Errorf("failed to delete image from GCS: %w", err)
	}

	log.Printf("Successfully deleted image with ID: %d", imageID)
	return nil
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

func (repo *Repository) GetLikeSearchImage(Hashtag string, CurrentID, LikeNum int) ([]models.Image, error) {
	var postedImages []models.PostedImage

	query := repo.db.
		Preload(clause.Associations).
		Table("posted_images").
		Select("posted_images.*, COUNT(posted_image_likes.user_id) as likes_count").
		Joins("LEFT JOIN posted_image_likes ON posted_image_likes.posted_image_id = posted_images.id").
		Joins("JOIN posted_image_hashtags ON posted_image_hashtags.posted_image_id = posted_images.id").
		Joins("JOIN hashtags ON posted_image_hashtags.hashtag_id = hashtags.id").
		Group("posted_images.id")

	// ハッシュタグ部分一致
	if Hashtag != "" {
		query = query.Where("hashtags.name LIKE ?", "%"+Hashtag+"%")
	}

	// ページネーション条件
	if LikeNum > -1 && CurrentID > -1 {
		query = query.Having(
			repo.db.
				Where("likes_count < ?", LikeNum).
				Or("likes_count = ? AND posted_images.id < ?", LikeNum, CurrentID),
		)
	}

	err := query.
		Order("likes_count DESC").
		Order("posted_images.id DESC").
		Limit(3).
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

func (repo *Repository) GetSearchImage(Hashtag string, CurrentID int) ([]models.Image, error) {
	var images []models.PostedImage
	query := repo.db.
		Preload(clause.Associations).
		Joins("JOIN posted_image_hashtags ON posted_image_hashtags.posted_image_id = posted_images.id").
		Joins("JOIN hashtags ON posted_image_hashtags.hashtag_id = hashtags.id")

	if Hashtag != "" {
		query = query.Where("hashtags.name LIKE ?", "%"+Hashtag+"%")
	}

	if CurrentID > 0 {
		query = query.Where("posted_images.id < ?", uint(CurrentID))
	}

	err := query.
		Group("posted_images.id").
		Order("posted_images.id DESC").
		Limit(3).
		Find(&images).Error

	if err != nil {
		return nil, err
	}

	var result []models.Image
	for _, img := range images {
		result = append(result, models.Image{
			ID:  img.ID,
			URL: img.URL,
		})
	}

	return result, nil
}

func (repo *Repository) GetLikeImages(CurrentID, LikeNum int) ([]models.Image, error) {
	var postedImages []models.PostedImage

	query := repo.db.
		Preload(clause.Associations).
		Table("posted_images").
		Select("posted_images.*, COUNT(posted_image_likes.user_id) as likes_count").
		Joins("LEFT JOIN posted_image_likes ON posted_image_likes.posted_image_id = posted_images.id").
		Group("posted_images.id")

	// ページネーション条件
	if LikeNum > -1 && CurrentID > -1 {
		query = query.Where(
			repo.db.
				Where("likes_count < ?", LikeNum).
				Or("likes_count = ? AND posted_images.id < ?", LikeNum, CurrentID),
		)
	}

	// 並び替え: いいね数の多い順 → IDの降順（安定した順序）
	err := query.
		Order("likes_count DESC").
		Order("posted_images.id DESC").
		Limit(3).
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

func (repo *Repository) GetRecentImages(CurrentID int) ([]models.Image, error) {
	var postedImages []models.PostedImage

	query := repo.db.
		Preload(clause.Associations).
		Order("posted_images.id DESC"). // 新しい順にするため ID の降順
		Limit(3)

	// ページネーション条件
	if CurrentID > 0 {
		query = query.Where("posted_images.id < ?", CurrentID)
	}

	err := query.Find(&postedImages).Error
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
