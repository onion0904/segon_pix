package repositories

import (
	"PixApp/models"
)

// 与えられたidのユーザー情報を返す
func (repo Repository) UserInfo (id uint) (models.User,error) {
	user := models.User{}
    if err := repo.db.First(&user, id).Error; err != nil {
        return user, err
    }
    return user, nil
}

// 与えられたハッシュタグの部分一致の画像のスライスを返す
func (repo Repository) ImageList (Qhashtag string) ([]models.PostedImage,error) {
	postimage := []models.PostedImage{}
	query := repo.db.Find(postimage)
	if Qhashtag != "" {
        query = query.Where("Hashtags LIKE ?", "%"+Qhashtag+"%")
    }

    if err := query.Find(&postimage).Error; err != nil {
        return nil, err
    }
    return postimage, nil
}