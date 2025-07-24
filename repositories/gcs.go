package repositories

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"time"

	"PixApp/models"
	"PixApp/util"
)

// UploadImageToGCS は画像を Google Cloud Storage にアップロードし、そのURLを返す。
func (repo *Repository) UploadImageToGCS(ctx context.Context, file io.Reader, filename string) (string, string, error) {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	if bucketName == "" {
		return "", "", fmt.Errorf("GCS bucket name is not set")
	}
	objectName := fmt.Sprintf("%v-%v", time.Now().Unix(), filename)
	object := repo.gcsClient.Bucket(bucketName).Object(objectName)
	wc := object.NewWriter(ctx)
	// ファイルの書き込み
	if _, err := io.Copy(wc, file); err != nil {
		wc.Close()
		return "", "", err
	}
	if err := wc.Close(); err != nil {
		return "", "", err
	}

	baseURL := "https://storage.googleapis.com/"
	url := fmt.Sprintf("%s%s/%s", baseURL, bucketName, objectName)
	return url, objectName, nil
}

// ensureHashtag は指定された名前のハッシュタグを取得、または存在しない場合は作成する。
func (repo *Repository) ensureHashtag(tx *gorm.DB, name string) (models.Hashtag, error) {
	var hashtag models.Hashtag
	err := tx.Where("Name = ?", name).First(&hashtag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ハッシュタグが存在しない場合は作成
			//この時、エラーが出るが気にしなくてよい
			hashtag = models.Hashtag{Name: name}
			if err := tx.Create(&hashtag).Error; err != nil {
				return models.Hashtag{}, fmt.Errorf("failed to create hashtag: %v", err)
			}
			return hashtag, nil
		}
		return hashtag, nil
	}
	return hashtag, err
}

func (repo *Repository) DeleteImageFromGCS(w io.Writer, ctx context.Context, objectName, bucketName string) error {
	if bucketName == "" {
		log.Printf("GCS bucket name is not set in environment variables")
		return fmt.Errorf("GCS bucket name is not set")
	}

	log.Printf("Attempting to delete object: %s from bucket: %s", objectName, bucketName)

	o := repo.gcsClient.Bucket(bucketName).Object(objectName)

	attrs, err := o.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %w", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %w", objectName, err)
	}

	fmt.Fprintf(w, "Blob %v deleted.\n", objectName)
	return nil
}

// ファイルを本番バケットから一時保管バケットに移動します。
func (repo *Repository) MoveImageToTemp(bucketName, subBucketName string, ctx context.Context, objectName string) (string, error) {
	// 移動元と移動先のGCSオブジェクトハンドルを取得
	src := repo.gcsClient.Bucket(bucketName).Object(objectName)
	dst := repo.gcsClient.Bucket(subBucketName).Object(objectName)

	// 移動先のオブジェクト(dst)に、移動元の内容(src)をコピーする
	copier := dst.CopierFrom(src)

	if _, err := copier.Run(ctx); err != nil {
		return "", fmt.Errorf("GCSオブジェクトのコピーに失敗: %w", err)
	}

	// コピーが成功したら、移動元のオブジェクトを削除する
	if err := src.Delete(ctx); err != nil {
		// コピーは成功したが削除に失敗した場合
		// コピー元GCS画像削除のリトライ処理
		const maxRetries = 5
		var gcsSrcDeleteErr error
		for attempt := 1; attempt <= maxRetries; attempt++ {
			gcsSrcDeleteErr = src.Delete(ctx)
			if gcsSrcDeleteErr == nil {
				// 削除に成功
				log.Printf("コピー元GCSオブジェクトのクリーンアップに成功しました。ObjectName: %s", objectName)
				// 成功したので、元のDBエラーを返して終了
				return objectName, nil
			}
			// 削除に失敗
			log.Printf("コピー元GCSオブジェクトのクリーンアップに失敗しました (試行 %d/%d)。エラー: %v", attempt, maxRetries, gcsSrcDeleteErr)
			time.Sleep(2 * time.Second) // 2秒待ってからリトライ
		}

		// リトライがすべて失敗した場合
		// コピー先元GCS画像削除のリトライ処理
		if gcsSrcDeleteErr != nil {
			const maxRetries = 5
			var gcsDstDeleteErr error
			for attempt := 1; attempt <= maxRetries; attempt++ {
				gcsDstDeleteErr = dst.Delete(ctx)
				if gcsDstDeleteErr == nil {
					// 削除に成功
					log.Printf("コピー先GCSオブジェクトのクリーンアップに成功しました。ObjectName: %s", objectName)
					// 成功したので、元のDBエラーを返して終了
					return "", fmt.Errorf("GCSオブジェクトのコピーに失敗: %w", err)
				}
				// 削除に失敗
				log.Printf("コピー先GCSオブジェクトのクリーンアップに失敗しました (試行 %d/%d)。エラー: %v", attempt, maxRetries, gcsDstDeleteErr)
				time.Sleep(2 * time.Second) // 2秒待ってからリトライ
			}

			if gcsDstDeleteErr != nil {
				// エラーメッセージを作成
				errorMessage := fmt.Sprintf("[本番の場所からの画像の保存先の移動] コピー先とコピー元のGCSオブジェクトが存在しています。コピー先の削除を手動で行ってください。 ObjectName: %s 最終エラー: 1:%v 2:%v", objectName, gcsSrcDeleteErr, gcsDstDeleteErr)

				// ログに致命的なエラーとして記録
				log.Println(errorMessage)

				// DBに失敗記録を残すための関数呼び出し
				err = repo.logFailToDB(objectName, gcsDstDeleteErr.Error())
				if err != nil {
					errorMessage = fmt.Sprintf("%s DBにログを残せませんでした: %v", errorMessage, err)
				}

				//slackに送信
				util.SlackNoticeTransaction(errorMessage)

				return "", fmt.Errorf("%s", errorMessage)
			}
		}
	}
	// 移動先のオブジェクト名を返す
	return objectName, nil
}

// ファイルを本番バケットから一時保管バケットに移動します。
func (repo *Repository) MoveTempToImage(subBucketName, bucketName string, ctx context.Context, objectName string) (string, error) {
	// 移動元と移動先のGCSオブジェクトハンドルを取得
	src := repo.gcsClient.Bucket(subBucketName).Object(objectName)
	dst := repo.gcsClient.Bucket(bucketName).Object(objectName)

	// 移動先のオブジェクト(dst)に、移動元の内容(src)をコピーする
	copier := dst.CopierFrom(src)

	// コピーが失敗したら、コピーをする
	if _, err := copier.Run(ctx); err != nil {
		const maxRetries = 5
		var copyErr error
		for attempt := 1; attempt <= maxRetries; attempt++ {
			_,copyErr = copier.Run(ctx)
			if copyErr == nil {
				// 削除に成功
				log.Printf("コピーに成功しました。ObjectName: %s", objectName)
				// 成功したので、元のDBエラーを返して終了
				return objectName, nil
			}
			// 削除に失敗
			log.Printf("コピーに失敗しました (試行 %d/%d)。エラー: %v", attempt, maxRetries, copyErr)
			time.Sleep(2 * time.Second) // 2秒待ってからリトライ
		}
		if copyErr != nil {
			errorMessage := fmt.Sprintf("[一時的な場所からの画像の保存先のコピー] コピーができませんでした。本番の場所に画像を移動させてください。手動での対応が必要です。 ObjectName: %s 最終エラー: %v", objectName, copyErr)

			// ログに致命的なエラーとして記録
			log.Println(errorMessage)

			// DBに失敗記録を残すための関数呼び出し
			err = repo.logFailToDB(objectName, err.Error())
			if err != nil {
				errorMessage = fmt.Sprintf("%s DBにログを残せませんでした: %v", errorMessage, err)
			}

			//slackに送信
			util.SlackNoticeTransaction(errorMessage)

			return "", fmt.Errorf("%s", errorMessage)
		}
	}

	// コピーが成功したら、移動元のオブジェクトを削除する
	if err := src.Delete(ctx); err != nil {
		// コピーは成功したが削除に失敗した場合
		// コピー元GCS画像削除のリトライ処理
		const maxRetries = 5
		var gcsSrcDeleteErr error
		for attempt := 1; attempt <= maxRetries; attempt++ {
			gcsSrcDeleteErr = src.Delete(ctx)
			if gcsSrcDeleteErr == nil {
				// 削除に成功
				log.Printf("コピー元GCSオブジェクトのクリーンアップに成功しました。ObjectName: %s", objectName)
				// 成功したので、元のDBエラーを返して終了
				return objectName, nil
			}
			// 削除に失敗
			log.Printf("コピー元GCSオブジェクトのクリーンアップに失敗しました (試行 %d/%d)。エラー: %v", attempt, maxRetries, gcsSrcDeleteErr)
			time.Sleep(2 * time.Second) // 2秒待ってからリトライ
		}

		if gcsSrcDeleteErr != nil {
			// エラーメッセージを作成
			errorMessage := fmt.Sprintf("[一時的な場所からの移動の画像保存元の削除] コピー先とコピー元のGCSオブジェクトが存在しています。一時的な場所から削除してください。手動での対応が必要です。 ObjectName: %s 最終エラー: %v", objectName, gcsSrcDeleteErr)

			// ログに致命的なエラーとして記録
			log.Println(errorMessage)

			// DBに失敗記録を残すための関数呼び出し
			err = repo.logFailToDB(objectName, gcsSrcDeleteErr.Error())
			if err != nil {
				errorMessage = fmt.Sprintf("%s DBにログを残せませんでした: %v", errorMessage, err)
			}

			//slackに送信
			util.SlackNoticeTransaction(errorMessage)

			return "", fmt.Errorf("%s", errorMessage)
		}
	}
	// 移動先のオブジェクト名を返す
	return objectName, nil
}
