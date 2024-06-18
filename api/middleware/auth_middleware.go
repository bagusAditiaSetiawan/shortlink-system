package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
	exception "shortlink-system/api/exceptions"
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
		panic(exception.NewUnauthorizedRequestException(languages.MALFORMED_JWT))
	}
	panic(exception.NewUnauthorizedRequestException(languages.EXPIRED_JWT))
}
