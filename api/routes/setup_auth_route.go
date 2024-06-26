package routes

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/api/handler"
)

func SetupAuthRoute(route fiber.Router, protectedMiddleware fiber.Handler, authHandler handler.AuthHandler) {
	route.Post("/auth/signup", authHandler.SignUp)
	route.Post("/auth/signin", authHandler.SignIn)
}
