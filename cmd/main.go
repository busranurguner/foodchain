package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/busranurguner/foodchain/docs"
	"github.com/busranurguner/foodchain/pkg/auth"
	"github.com/busranurguner/foodchain/pkg/db"
	"github.com/busranurguner/foodchain/pkg/logger"
	"github.com/busranurguner/foodchain/pkg/middlewares"
	"github.com/busranurguner/foodchain/pkg/user"
	"github.com/gofiber/fiber/v2"

	log "github.com/gofiber/fiber/v2/middleware/logger"
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

	//Run db
	database := db.ConfigDB()

	//Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//Logging
	zapLogger, _ := logger.NewZapLogger()
	logger.L = zapLogger

	app.Use(log.New())

	//Grouping
	v1 := app.Group("/v1")

	authRepo := auth.NewRepository(database)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	v1.Post("/signup", authHandler.SignUp)
	v1.Post("/login", authHandler.Login)
	v1.Post("/refresh", authHandler.Refresh)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mykey"),
		AuthScheme: "Bearer",
	}))

	userRepo := user.NewRepository(database)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	v1.Get("/user", middlewares.BasicAuthHandler(), userHandler.GetAll)

	app.Listen(":3000")

	logger.L.Fatal(app.Listen(":3000"))
}
