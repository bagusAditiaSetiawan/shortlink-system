package handler

import (
	"github.com/gofiber/fiber/v2"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/helper"
)

type ProfileHandlerImpl struct {
	UserLoggedService auth.UserLoggedService
}

func NewProfileHandlerImpl(userLoggedService auth.UserLoggedService) *ProfileHandlerImpl {
	return &ProfileHandlerImpl{
		UserLoggedService: userLoggedService,
	}
}

func (handler *ProfileHandlerImpl) GetProfile(ctx *fiber.Ctx) error {
	profile := handler.UserLoggedService.GetUserLoggedPayload(ctx)
	return ctx.JSON(helper.ToWebResponse(profile))
}
