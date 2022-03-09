package controllers

import (
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func SignUp(ctx *fiber.Ctx) error {

	return nil
}

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
	atoken, rtoken, err := Token(foundUser.Username, foundUser.Password, foundUser.Role)
	if err != nil {
		return err
	}

	//her login oldugunda yeni bir refresh token olmalı mı?
	update := bson.M{"refresh": rtoken}
	userCollection.UpdateOne(ctx.Context(), bson.M{"_id": foundUser.ID}, bson.M{"$set": update})

	return ctx.JSON(fiber.Map{
		"access":  atoken,
		"refresh": rtoken,
	})
}

func RefreshToken(ctx *fiber.Ctx) error {

	var token models.Token
	var fuser models.User

	err := ctx.BodyParser(&token)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	err = userCollection.FindOne(ctx.Context(), bson.M{"refresh": token.RefreshToken}).Decode(&fuser)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not Found",
		})
	}
	atoken, rtoken, err := Token(fuser.Username, fuser.Password, fuser.Role)
	if err != nil {
		return err
	}

	update := bson.M{"refresh": rtoken}
	userCollection.UpdateOne(ctx.Context(), bson.M{"_id": fuser.ID}, bson.M{"$set": update})

	return ctx.JSON(fiber.Map{
		"access":  atoken,
		"refresh": rtoken,
	})
}
