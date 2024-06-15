package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/api/handler"
	"shortlink-system/api/middleware"
	"shortlink-system/api/routes"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/database"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/jwt"
)

func main() {
	err := godotenv.Load()
	helper.IfErrorHandler(err)
	app := initializedApp()
	validate := validator.New()
	db := database.InitializedDatabase()
	InitializedService(app, validate, db)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

func InitializedService(app *fiber.App, validate *validator.Validate, db *gorm.DB) {
	userService := auth.InitializedAuthService(db, validate)

	jwtService := jwt.InitializedJwt()
	authHandler := handler.NewAuthHandler(userService, jwtService)

	grouped := routes.SetupRouteApi(app)
	routes.SetupAuthRoute(grouped, authHandler)

	protectedMiddleware := middleware.Protected()
	userLoggedService := auth.NewUserLoggedService()
	profileHandler := handler.NewProfileHandlerImpl(userLoggedService)
	routes.SetupProfileRoute(grouped, protectedMiddleware, profileHandler)
}

func initializedApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandlerException,
	})
	app.Use(recover.New())
	return app
}
