package controllers

import (
	_ "github.com/busranurguner/foodchain/docs"
	"github.com/busranurguner/foodchain/pkg/db"
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.GetCollection(db.DB, "User")

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
	role := claims["role"].(string)

	if role != "admin" {
		return ctx.SendStatus(401)
	}

	var users []models.User
	cursor, err := userCollection.Find(ctx.Context(), bson.M{}) // bson.M gelen verilerin sırası önemsiz ise kullanılır.
	if err != nil {
		return nil
	}

	for cursor.Next(ctx.Context()) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			return ctx.SendStatus(500)
		}
		users = append(users, user)
	}
	return ctx.JSON(users)

}
