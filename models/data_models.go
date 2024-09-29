package models

import(
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name         string         // ユーザー名
    Profile      string         // プロフィールメッセージ
    Icon         string         // プロフィール画像のURL
    HeaderImage  string         // ヘッダー画像のURL
    Email        *string        // メールアドレス（任意）
    Age          *uint8         // 年齢（任意）
    PostedImages []PostedImage  // ユーザーが投稿した画像
    LikedImages  []LikedImage   // ユーザーがいいねした画像
    Follows      []*User        `gorm:"many2many:user_follows;"` // フォロー関係（多対多）
}

type PostedImage struct {
    gorm.Model
    URL    string `gorm:"not null"` // 画像のURL
    UserID uint   `gorm:"not null"` // 投稿者のユーザーID
}

type LikedImage struct {
    gorm.Model
    URL    string `gorm:"not null"` // 画像のURL
    UserID uint   `gorm:"not null"` // いいねしたユーザーのID
}
