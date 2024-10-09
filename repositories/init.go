package repositories

import (
    "gorm.io/gorm"
	"cloud.google.com/go/storage"
	"context"
    "github.com/joho/godotenv"
    "fmt"
    "log"
    "google.golang.org/api/option"
)

type Repository struct {
	db *gorm.DB
	gcsClient *storage.Client
}

func NewRepository(db *gorm.DB) (*Repository, error) {
    // .env ファイルの読み込み
    log.Println("Attempting to load .env file...")
    err := godotenv.Load()
    if err != nil {
        log.Printf("Error loading .env file: %v", err)
        return nil, fmt.Errorf("failed to load .env file: %w", err)
    }
    log.Println(".env file loaded successfully")

    // Google Cloud Storage クライアントの作成
    log.Println("Attempting to create Google Cloud Storage client...")
    ctx := context.Background()
    client, err := storage.NewClient(ctx, option.WithCredentialsFile("/app/myapp-437007-bdde37cabb9b.json"))
    if err != nil {
        log.Printf("Error creating GCS client: %v", err)
        return nil, fmt.Errorf("failed to create GCS client: %w", err)
    }
    log.Println("Google Cloud Storage client created successfully")

    return &Repository{db: db, gcsClient: client}, nil
}



// Closeメソッドを追加して、クライアントのリソースを適切に解放
func (r *Repository) Close() error {
    return r.gcsClient.Close()
}

