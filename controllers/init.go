package controllers

import (
    "gorm.io/gorm"
	"sync"
    "strconv"
    "log"
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


func uintID(userID string) uint{
    uintID64, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        log.Printf("Invalid user ID format: %v", err)
        return 0
    }
    uintID := uint(uintID64)
    return uintID
}