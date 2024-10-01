package controllers

import (
    "gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

// コンストラクタ
func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}