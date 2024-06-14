package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/api/handler"
	"shortlink-system/api/routes"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/database"
)

func main() {
	app := initializedApp()
	validate := validator.New()
	db := database.InitializedDatabase()
	userService := auth.InitializedAuthService(db, validate)

	authHandler := handler.NewAuthHandler(userService)

	grouped := routes.SetupRouteApi(app)
	routes.SetupAuthRoute(grouped, authHandler)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

func initializedApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandlerException,
	})
	return app
}
