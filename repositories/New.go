package repositories

import (
    "gorm.io/gorm"
	"cloud.google.com/go/storage"
	"context"
    "fmt"
    "google.golang.org/api/option"
)

type Repository struct {
    db        *gorm.DB
    gcsClient *storage.Client
}

func NewRepository(db *gorm.DB) (*Repository, error) {
    ctx := context.Background()
    client, err := storage.NewClient(ctx, option.WithCredentialsFile("/app/myapp-437007-bdde37cabb9b.json"))
    if err != nil {
        return nil, fmt.Errorf("failed to create GCS client: %w", err)
    }

    return &Repository{db: db, gcsClient: client}, nil
}

// Closeメソッドを追加して、クライアントのリソースを適切に解放
func (r *Repository) Close() error {
    return r.gcsClient.Close()
}

func (r *Repository) DB() *gorm.DB {
    return r.db
}


