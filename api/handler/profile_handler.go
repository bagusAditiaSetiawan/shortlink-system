package handler

import "github.com/gofiber/fiber/v2"

type ProfileHandler interface {
	GetProfile(ctx *fiber.Ctx) error
}
