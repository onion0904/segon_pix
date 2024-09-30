package repositories

import (
    "gorm.io/gorm"
	"cloud.google.com/go/storage"
	"context"
)

type Repository struct {
	db *gorm.DB
	gcsClient *storage.Client
}

func NewRepository(db *gorm.DB) (*Repository, error) {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        return nil, err
    }
    return &Repository{db: db, gcsClient: client}, nil
}