package repositories

import (
    "PixApp/models"
    "errors"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)


func (repo *Repository) ExistUser(email string, password string) (bool,error) {
    var user models.User
    // Emailでユーザーを検索
    if err := repo.db.
        Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false,nil
        }
        return false,err
    }

    // パスワードの照合
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return false,errors.New("invalid password")
    }

    return true,nil
}