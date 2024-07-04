package auth

import "github.com/gofiber/fiber/v2"

type UserLoggedPayload struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoggedService interface {
	GetUserLoggedPayload(ctx *fiber.Ctx) UserLoggedPayload
}
