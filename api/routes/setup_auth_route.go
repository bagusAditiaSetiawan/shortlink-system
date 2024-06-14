package routes

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/api/handler"
)

func SetupAuthRoute(route fiber.Router, authHandler handler.AuthHandler) {
	route.Post("/auth/signup", authHandler.Signup)
}
