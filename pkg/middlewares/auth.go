package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func BasicAuthHandler() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		u := c.Locals("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "admin" {
			return c.SendStatus(401)
		}
		return c.Next()
	}
}
