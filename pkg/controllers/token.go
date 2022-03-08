package controllers

import (
	"time"

	_ "github.com/busranurguner/foodchain/docs"
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("mykey") //configten al.

func Token() (map[string]string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "busra",
		"admin": true,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}
	// Create access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	// Create refresh token
	rclaims := jwt.MapClaims{
		"uuid": 1,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)

	rtokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": rtokenString,
	}, nil
}

func RefreshToken(ctx *fiber.Ctx) error {
	var rtoken models.Token

	err := ctx.BodyParser(&rtoken)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	token, err := Token()
	if err != nil {
		return nil
	}
	return ctx.JSON(token)
}
