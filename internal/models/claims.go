package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
