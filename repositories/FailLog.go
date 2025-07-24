package repositories

import (
	"fmt"

	"PixApp/models"
)

func (repo *Repository) logFailToDB(objectName, gcsDeleteErr string) error {
	logFailDB := &models.LogFailDB{
		ObjectName: objectName,
		Error:      gcsDeleteErr,
	}

	// 失敗ログを追加
	if err := repo.db.Create(logFailDB).Error; err != nil {
		return fmt.Errorf("failed to add fail gcs log to the database: %w", err)
	}
	return nil
}