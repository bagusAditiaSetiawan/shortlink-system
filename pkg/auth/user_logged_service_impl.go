package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserLoggedServiceImpl struct{}

func NewUserLoggedService() *UserLoggedServiceImpl {
	return &UserLoggedServiceImpl{}
}

func (service *UserLoggedServiceImpl) GetUserLoggedPayload(ctx *fiber.Ctx) UserLoggedPayload {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	userLoggedPayload := UserLoggedPayload{
		ID:       uint(id),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
	}
	return userLoggedPayload
}
