package shorted_link

import (
	"gorm.io/gorm"
	"shortlink-system/pkg/entities"
)

type ShortedLinkRepository interface {
	Create(db *gorm.DB, req entities.ShortedLink) entities.ShortedLink
	FindByOriginalLink(db *gorm.DB, link string) (entities.ShortedLink, error)
	FindByShortedLink(db *gorm.DB, link string) (entities.ShortedLink, error)
}
