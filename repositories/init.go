package repositories

import (
    "gorm.io/gorm"
	"cloud.google.com/go/storage"
	"context"
    "github.com/joho/godotenv"
    "log"
)

type Repository struct {
	db *gorm.DB
	gcsClient *storage.Client
}

func NewRepository(db *gorm.DB) (*Repository, error) {
    // .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        return nil, err
    }
    return &Repository{db: db, gcsClient: client}, nil
}