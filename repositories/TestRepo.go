package repositories

import (
    _"gorm.io/driver/mysql"
    "gorm.io/gorm"
	"PixApp/models"
	"gorm.io/gorm/logger"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
    return &TodoRepository{db: db}
}

func (repo TodoRepository) Add (model *models.Todo) error {
	if err := repo.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (repo TodoRepository) List (Qtitle string) ([]models.Todo,error) {
	todos := []models.Todo{}
	query := repo.db.Select("ID","Title","CreatedAt","UpdatedAt","Priority","Status")
	if Qtitle != "" {
        query = query.Where("Title LIKE ?", "%"+Qtitle+"%")
    }

    if err := query.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

func (repo TodoRepository) Update (newTodo *models.Todo) error {
	repo.db.Logger = repo.db.Logger.LogMode(logger.Info)
	
	if err := repo.db.Model(&models.Todo{}).Where("ID = ?", newTodo.ID).Updates(newTodo).Error; err != nil {
		return err
	}
	return nil
}

func (repo TodoRepository) Delete (id int) error {
	if err := repo.db.Delete(&models.Todo{},id).Error; err != nil {
		return err
	}
	return nil
}