package routes

import "github.com/gofiber/fiber/v2"

func SetupRouteApi(app *fiber.App) fiber.Router {
	grouped := app.Group("/api")
	return grouped
}
