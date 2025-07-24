package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"

	"PixApp/models"
	"PixApp/repositories"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 環境変数の取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")

	// データベース接続の設定

	dsn := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?parseTime=true", dbUser, dbPassword, os.Getenv("DB_PORT"), dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo, err := repositories.NewRepository(db)
	if err != nil {
		log.Printf("Failed to create repository: %v", err)
	}

	var logs []models.LogFailDB
	for {
		if err := db.Find(&logs).Error; err != nil {
			// エラー処理
			log.Printf("DB取得エラー: %v", err)
		}

		for i := range logs {
			switch {
			case strings.HasPrefix(logs[i].Error, "[画像の追加機能]"):
				deleteImageFromGCS(repo, logs[i], os.Getenv("GCS_BUCKET_NAME"))
			case strings.HasPrefix(logs[i].Error, "[画像ファイル一時的な場所から削除]"):
				deleteImageFromGCS(repo, logs[i], os.Getenv("GCS_SUB_BUCKET_NAME"))
			case strings.HasPrefix(logs[i].Error, "[本番の場所からの画像の保存先の移動]"):
				deleteImageFromGCS(repo, logs[i], os.Getenv("GCS_SUB_BUCKET_NAME"))
			case strings.HasPrefix(logs[i].Error, "[一時的な場所からの移動の画像保存元の削除]"):
				deleteImageFromGCS(repo, logs[i], os.Getenv("GCS_SUB_BUCKET_NAME"))
			}
		}
		time.Sleep(24 * time.Hour)
	}
}

func deleteLogFromDB(db *gorm.DB, objectName string) error {
	if err := db.Where("object_name = ?", objectName).Delete(&models.LogFailDB{}).Error; err != nil {
		return err
	}
	return nil
}

func deleteImageFromGCS(repo *repositories.Repository, logFail models.LogFailDB, bucketName string) {
	err := repo.DeleteImageFromGCS(os.Stdout, context.Background(), logFail.ObjectName, bucketName)
	if err == nil {
		if err = deleteLogFromDB(repo.DB(), logFail.ObjectName); err != nil {
			log.Fatalf("failed to deleteLogFromDB: %v", err)
		}
	}
}
