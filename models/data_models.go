package models

import (
    "gorm.io/gorm"
    "github.com/golang-jwt/jwt/v5"
)

type User struct {
    gorm.Model
    Name         string        `gorm:"not null;index"` // ユーザー名
    Description  string        // プロフィールメッセージ
    Icon         string        // プロフィール画像のURL
    HeaderImage  string        // ヘッダー画像のURL
    Email        string        `gorm:"type:varchar(100);unique"` // メールアドレス（一意のメール形式）
    Password     string        `gorm:"type:varchar(100);unique"`// パスワード
    Birthday     int           // 誕生日(20241009=2024年10月9日)
    PostedImages []PostedImage `gorm:"foreignKey:UserID"`// ユーザーが投稿した画像
    LikedImages  []PostedImage `gorm:"many2many:posted_image_likes;constraint:OnDelete:CASCADE"`
    Follows      []User        `gorm:"many2many:user_follows;joinForeignKey:FollowerID;JoinReferences:FollowingID"` // フォローしているユーザー
    Followers    []User        `gorm:"many2many:user_follows;joinForeignKey:FollowingID;JoinReferences:FollowerID"` // フォローされているユーザー
}

type PostedImage struct {
    gorm.Model
    URL         string    `gorm:"not null"`        
    UserID      uint      `gorm:"not null"` // 画像を投稿した人のID       
    PostUser    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`  // 投稿者が削除されたら関連する画像も削除
    ObjectName  string    `gorm:"not null;index"`  // バックエンドの処理で使用
    Likes       []User    `gorm:"many2many:posted_image_likes;constraint:OnDelete:CASCADE"` // いいねしたユーザー
    Comments    []Comment `gorm:"constraint:OnDelete:CASCADE"` // コメントが削除されたときにリレーションを更新
    Hashtags    []Hashtag `gorm:"many2many:posted_image_hashtags;"` 
}

type Comment struct {
    gorm.Model
    PostedImageID uint   `gorm:"not null"` // コメント元のPostedImageのID
    UserID        uint   `gorm:"not null"` // コメントしたユーザーのID (外部キー)
    Message       string // コメント内容
}



type Hashtag struct {
    gorm.Model
    Name         string        `gorm:"unique;not null"` // ハッシュタグ名
    PostedImages []PostedImage `gorm:"many2many:posted_image_hashtags;"` // 同じハッシュタグを持つ画像
}






type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

//repositoriesのPostedImageの中で画像を返す時に使用
type Image struct{
    ID       uint
    URL      string
}
