package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name         string        `gorm:"not null;index"` // ユーザー名
    Description  string        // プロフィールメッセージ
    Icon         string        // プロフィール画像のURL
    HeaderImage  string        // ヘッダー画像のURL
    Email        string        `gorm:"type:varchar(100);unique"` // メールアドレス（任意、一意のメール形式）
    Birthday     int           // 誕生日
    PostedImages []PostedImage `gorm:"foreignKey:UserID"`// ユーザーが投稿した画像
    LikedImages  []PostedImage `gorm:"many2many:posted_image_likes;constraint:OnDelete:CASCADE"`
    Follows      []User        `gorm:"many2many:user_follows;joinForeignKey:FollowerID;JoinReferences:FollowingID"` // フォローしているユーザー
    Followers    []User        `gorm:"many2many:user_follows;joinForeignKey:FollowingID;JoinReferences:FollowerID"` // フォローされているユーザー
}

type PostedImage struct {
    gorm.Model
    URL         string    `gorm:"not null"`        
    UserID      uint      `gorm:"not null"` //画像を投稿した人のID       
    PostUser    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`  // 投稿者が削除されたら関連する画像も削除
    ObjectName  string    `gorm:"not null;index"`  // バックエンドの処理で使用
    Likes       []User    `gorm:"many2many:posted_image_likes;constraint:OnDelete:CASCADE"` // いいねしたユーザー
    Comments    []Comment 
    Hashtags    []Hashtag `gorm:"many2many:posted_image_hashtags;"` 
}


type Comment struct {
    gorm.Model
    PostedImageID uint   `gorm:"not null"` // コメント元のPostedImageのID
    UserID        uint   `gorm:"not null"` // コメントしたユーザーのID (外部キー)
    PostUser      User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // 外部キーを明示的に指定
    Message       string // コメント内容
}


type Hashtag struct {
    gorm.Model
    Name         string        `gorm:"unique;not null"` // ハッシュタグ名
    PostedImages []PostedImage `gorm:"many2many:posted_image_hashtags;"` // 同じハッシュタグを持つ画像
}
