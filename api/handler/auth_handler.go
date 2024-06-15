package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
}
