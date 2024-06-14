package handler

import (
	"github.com/gofiber/fiber/v2"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/languages"
)

type AuthHandlerImpl struct {
	AuthService auth.Service
}

func NewAuthHandler(userService auth.Service) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		AuthService: userService,
	}
}

func (handler *AuthHandlerImpl) Signup(ctx *fiber.Ctx) error {
	req := new(auth.SignUpRequest)
	if err := ctx.BodyParser(req); err != nil {
		panic(exception.NewBadRequestException(languages.MALFORMED))
	}
	user := handler.AuthService.SignUp(req)
	return ctx.Status(fiber.StatusCreated).JSON(helper.ToWebResponse(user))
}
