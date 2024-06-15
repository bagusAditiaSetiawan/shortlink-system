package shorted_link

import "shortlink-system/pkg/entities"

type ShortLinkService interface {
	Handler(params *CreateShortedLink, user entities.User) entities.ShortedLink
	GetExistShortLink(linkOriginal string) (string, error)
	GetExistOriginalLink(linkOriginal string) (string, error)
	SaveToRedis(link entities.ShortedLink)
}
