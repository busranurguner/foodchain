package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/busranurguner/foodchain/docs"
	"github.com/busranurguner/foodchain/pkg/controllers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// @title User API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	// Login route
	app.Post("/login", controllers.Login)

	app.Get("/swagger/*", swagger.HandlerDefault)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/user", controllers.GetUser)

	app.Listen(":3000")
}
