package shorted_link

import (
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/entities"
)

type ShortLinkService interface {
	CreateShortedLink(params *CreateShortedLink, user auth.UserLoggedPayload) entities.ShortedLink
	GetExistOriginalLink(generatedLink string) (string, error)
	SaveToRedis(link entities.ShortedLink)
	UpdateAccessed(link string) entities.ShortedLink
	PaginateShortLink(req *PaginateShortedLink) ([]entities.ShortedLink, int)
}
