package routes

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/api/handler"
)

func SetupShortedLinkRoute(route fiber.Router, protected fiber.Handler, handler handler.ShortedLinkHandler) {
	route.Post("/shorted-link", protected, handler.CreateShortedLink)
	route.Post("/shorted-link/list", protected, handler.PaginateLink)
}
