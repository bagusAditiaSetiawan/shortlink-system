package handler

import (
	"github.com/gofiber/fiber/v2"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/api/presenter"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/languages"
	"shortlink-system/pkg/shorted_link"
)

type ShortedLinkHandlerImpl struct {
	UserLoggedService  auth.UserLoggedService
	ShortedLinkService shorted_link.ShortLinkService
}

func NewShortedLinkHandlerImpl(userLoggedService auth.UserLoggedService, shortedLinkService shorted_link.ShortLinkService) *ShortedLinkHandlerImpl {
	return &ShortedLinkHandlerImpl{
		UserLoggedService:  userLoggedService,
		ShortedLinkService: shortedLinkService,
	}
}

func (handler *ShortedLinkHandlerImpl) CreateShortedLink(ctx *fiber.Ctx) error {
	req := new(shorted_link.CreateShortedLink)
	if err := ctx.BodyParser(req); err != nil {
		panic(exception.NewBadRequestException(languages.MALFORMED))
	}
	user := handler.UserLoggedService.GetUserLoggedPayload(ctx)
	shortedLink := handler.ShortedLinkService.CreateShortedLink(req, user)

	return ctx.Status(fiber.StatusCreated).JSON(helper.ToWebResponse(fiber.Map{
		"shorted_link": shortedLink.LinkShorted,
	}))
}

func (handler *ShortedLinkHandlerImpl) RedirectLink(ctx *fiber.Ctx) error {
	link := ctx.Params("link")
	originalLink, err := handler.ShortedLinkService.GetExistOriginalLink(link)
	if err != nil {
		panic(exception.NewNotFoundHandler(languages.LINK_NOT_FOUND))
	}
	handler.ShortedLinkService.UpdateAccessed(link)
	return ctx.Redirect(originalLink, 301)
}

func (handler *ShortedLinkHandlerImpl) PaginateLink(ctx *fiber.Ctx) error {
	req := new(shorted_link.PaginateShortedLink)
	if err := ctx.BodyParser(req); err != nil {
		panic(exception.NewBadRequestException(languages.MALFORMED))
	}
	user := handler.UserLoggedService.GetUserLoggedPayload(ctx)
	req.UserID = user.ID
	shortedLinks, total := handler.ShortedLinkService.PaginateShortLink(req)
	return ctx.JSON(helper.ToWebResponse(presenter.ToShortedLinkResponse(shortedLinks, total)))
}
