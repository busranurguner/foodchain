package controllers

import (
	"time"

	_ "github.com/busranurguner/foodchain/docs"
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

//var secretKey = []byte("mykey") //configten al.

// GetToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Param token body models.UserToken true "Create Token"
// @Success 200 {array} models.Token
// @Router /login [post]
func Login(ctx *fiber.Ctx) error {

	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	if user.Username != "busra" || user.Password != "123" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "busra",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{"token": tokenString})
}
