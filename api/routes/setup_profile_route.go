package routes

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/api/handler"
)

func SetupProfileRoute(route fiber.Router, protected fiber.Handler, profileHandler handler.ProfileHandler) {
	route.Get("/profile", protected, profileHandler.GetProfile)
}
