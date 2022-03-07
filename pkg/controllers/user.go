package controllers

import (
	_ "github.com/busranurguner/foodchain/docs"
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetUser method.
// @Description Get user.
// @Summary Get user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /user [get]
// @Security ApiKeyAuth
func GetUser(ctx *fiber.Ctx) error {

	u := ctx.Locals("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	role := claims["admin"].(bool)

	if !role {
		return ctx.SendStatus(401)
	}

	user := []models.User{
		{
			ID:       "1",
			Username: "busra",
			Role:     "admin",
			Password: "123",
		},
		{
			ID:       "2",
			Username: "ayse",
			Role:     "member",
			Password: "456",
		},
	}
	return ctx.JSON(user)
}
