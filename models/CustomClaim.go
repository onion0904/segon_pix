package models

import (
    "github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Email string `json:"email,omitempty"`
    UserID uint `json:"userid,omitempty"`
	jwt.RegisteredClaims
}