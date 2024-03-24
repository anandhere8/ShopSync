package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Token struct {
	Access_token string `json:"access_token"`
	Auth_type    string `json:"auth_type"`
	ExpiresAt    string `json:"after"`
	CreatedAt    string `json:"createdAt"`
	Issuer       string `json:"issuer"`
}
