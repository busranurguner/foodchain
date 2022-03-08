package controllers

import (
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

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
	var foundUser models.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot Parse JSON",
		})
	}
	err = userCollection.FindOne(ctx.Context(), bson.M{"username": user.Username, "password": user.Password}).Decode(&foundUser)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}
	token, err := Token()
	if err != nil {
		return err
	}
	return ctx.JSON(token)
}
