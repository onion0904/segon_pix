package repositories

import (
    "PixApp/models"
    "errors"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)


func (repo *Repository) ExistUser(email string, password string) (uint,error) {
    var user models.User
    // Emailでユーザーを検索
    if err := repo.db.
        Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return 0,nil
        }
        return 0,err
    }

    // パスワードの照合
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return 0,errors.New("invalid password")
    }

    return user.ID,nil
}