package repositories

import (
    "PixApp/models"
)


func (repo *Repository) ExistUser(email string, password string) error {
    var user models.User
    if err := repo.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
        return err 
    }
    return nil
}
