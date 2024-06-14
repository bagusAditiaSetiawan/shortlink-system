package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Signup(ctx *fiber.Ctx) error
}
