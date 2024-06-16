package handler

import "github.com/gofiber/fiber/v2"

type ShortedLinkHandler interface {
	CreateShortedLink(ctx *fiber.Ctx) error
	RedirectLink(ctx *fiber.Ctx) error
	PaginateLink(ctx *fiber.Ctx) error
}
