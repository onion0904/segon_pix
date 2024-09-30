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
    Birthday     *int           // 誕生日（任意）
    PostedImages []*PostedImage // ユーザーが投稿した画像
    LikedImages  []*PostedImage // ユーザーがいいねした画像
    Follows      []*User       `gorm:"many2many:user_follows;joinForeignKey:FollowerID;JoinReferences:FollowingID"` // フォローしているユーザー
    Followers    []*User       `gorm:"many2many:user_follows;joinForeignKey:FollowingID;JoinReferences:FollowerID"` // フォローされているユーザー
}

type PostedImage struct {
    gorm.Model
    URL         string    `gorm:"not null"`                    // 画像のURL
    UserID      uint      `gorm:"not null"`                    // 投稿者のユーザーID
    PostUser    User      `gorm:"not null"`                    //投稿したユーザーの情報
    Likes       []User    // いいねしたユーザー
    Comments    []Comment // コメント
    Hashtags    []Hashtag `gorm:"many2many:posted_image_hashtags;"` // ハッシュタグ
}

type Comment struct {
    gorm.Model
    PostedImageID uint   `gorm:"not null"` // コメント元のPostedImageのID
    UserID        uint   `gorm:"not null"` // コメントしたユーザーのID
    PostUser      User   `gorm:"not null"` // コメントしたユーザーの情報
    Message       string // コメント内容
}

type Hashtag struct {
    gorm.Model
    Name         string        `gorm:"unique;not null"` // ハッシュタグ名
    PostedImages []PostedImage `gorm:"many2many:posted_image_hashtags;"` // 同じハッシュタグを持つ画像
}
