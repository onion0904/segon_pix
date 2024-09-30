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
func (repo *Repository) ImageList(Qhashtag string) ([]models.PostedImage, error) {
	var images []models.PostedImage
	err := repo.db.
		Model(&models.PostedImage{}).
		Joins("JOIN posted_image_hashtags on posted_image_hashtags.posted_image_id = posted_images.id").
		Joins("JOIN hashtags on posted_image_hashtags.hashtag_id = hashtags.id").
		Where("hashtags.name LIKE ?", "%"+Qhashtag+"%").
		Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}