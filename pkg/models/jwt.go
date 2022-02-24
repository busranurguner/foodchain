package models

import "github.com/golang-jwt/jwt"

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
	Status       string `json:"status"`
}

type Claim struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
