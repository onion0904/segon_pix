package controllers

import (
    "gorm.io/gorm"
	"sync"
)

type Controller struct {
    db   *gorm.DB        // データベース接続
    auth *Auth           // Auth構造体のインスタンスを保持
}

// コンストラクタ
func NewController(db *gorm.DB) *Controller {
    return &Controller{
        db:   db,
        auth: NewAuth(),  // Authの初期化
    }
}


type Auth struct {
	VerificationCodes map[string]string
	CodeMutex         sync.Mutex
}

// コンストラクタ
func NewAuth() *Auth {
	return &Auth{
		VerificationCodes: make(map[string]string),
	}
}