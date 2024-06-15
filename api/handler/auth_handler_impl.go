package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/jwt"
	"shortlink-system/pkg/languages"
)

type AuthHandlerImpl struct {
	AuthService auth.Service
	JwtService  jwt.JwtService
}

func NewAuthHandler(userService auth.Service, jwtService jwt.JwtService) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		AuthService: userService,
		JwtService:  jwtService,
	}
}

func (handler *AuthHandlerImpl) SignUp(ctx *fiber.Ctx) error {
	req := new(auth.SignUpRequest)
	if err := ctx.BodyParser(req); err != nil {
		panic(exception.NewBadRequestException(languages.MALFORMED))
	}
	user := handler.AuthService.SignUp(req)
	return ctx.Status(fiber.StatusCreated).JSON(helper.ToWebResponse(user))
}
func (handler *AuthHandlerImpl) SignIn(ctx *fiber.Ctx) error {
	req := new(auth.SignInRequest)
	if err := ctx.BodyParser(req); err != nil {
		panic(exception.NewBadRequestException(languages.MALFORMED))
	}
	log.Info("Process signin request ", req.Username)
	user := handler.AuthService.SignIn(req)
	log.Info("process generate token ", user.ID)
	accessToken := handler.JwtService.Generate(&jwt.TokenPayload{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
	return ctx.Status(fiber.StatusCreated).JSON(helper.ToWebResponse(fiber.Map{
		"access_token": accessToken,
	}))
}
