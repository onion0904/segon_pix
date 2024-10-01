package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name         string        `gorm:"not null;index"` // ユーザー名
    Profile      *string        // プロフィールメッセージ
    Icon         *string        // プロフィール画像のURL
    HeaderImage  *string        // ヘッダー画像のURL
    Email        *string       `gorm:"type:varchar(100);unique"` // メールアドレス（任意、一意のメール形式）
    Birthday     int           // 誕生日
    PostedImages []*PostedImage // ユーザーが投稿した画像
    LikedImages  []*PostedImage // ユーザーがいいねした画像
    Follows      []*User       `gorm:"many2many:user_follows;joinForeignKey:FollowerID;JoinReferences:FollowingID"` // フォローしているユーザー
    Followers    []*User       `gorm:"many2many:user_follows;joinForeignKey:FollowingID;JoinReferences:FollowerID"` // フォローされているユーザー
}

type PostedImage struct {
    gorm.Model
    URL         string    `gorm:"not null"`        // 画像のURL
    UserID      uint      `gorm:"not null"`        // 投稿者のユーザーID (外部キー)
    PostUser    User      `gorm:"foreignKey:UserID"`  // 外部キーを明示的に指定
    Likes       []User    `gorm:"many2many:posted_image_likes;"` // いいねしたユーザー（多対多リレーション）
    Comments    []Comment // コメント
    Hashtags    []Hashtag `gorm:"many2many:posted_image_hashtags;"` // ハッシュタグ
}

type Comment struct {
    gorm.Model
    PostedImageID uint   `gorm:"not null"` // コメント元のPostedImageのID
    UserID        uint   `gorm:"not null"` // コメントしたユーザーのID (外部キー)
    PostUser      User   `gorm:"foreignKey:UserID"` // 外部キーを明示的に指定
    Message       string // コメント内容
}


type Hashtag struct {
    gorm.Model
    Name         string        `gorm:"unique;not null"` // ハッシュタグ名
    PostedImages []PostedImage `gorm:"many2many:posted_image_hashtags;"` // 同じハッシュタグを持つ画像
}
