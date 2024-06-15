package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
	"shortlink-system/pkg/languages"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("TOKEN_KEY"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"errors": []string{
					languages.MALFORMED_JWT,
				},
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"errors": []string{
			languages.EXPIRED_JWT,
		}})
}
