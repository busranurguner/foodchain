package controllers

import (
	"log"
	"time"

	_ "github.com/busranurguner/foodchain/docs"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("mykey") //configten al.

func Token(username string, password string, role string) (atoken string, rtoken string, err error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"role":     role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}
	// Create access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create refresh token
	rclaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)

	rtokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, rtokenString, err
}
