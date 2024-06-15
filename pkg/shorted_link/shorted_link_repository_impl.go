package shorted_link

import (
	"gorm.io/gorm"
	"shortlink-system/pkg/entities"
)

type ShortedLinkRepositoryImpl struct {
}

func (repository *ShortedLinkRepositoryImpl) Create(db *gorm.DB, req entities.ShortedLink) entities.ShortedLink {
	db.Create(&req)
	return req
}
func (repository *ShortedLinkRepositoryImpl) FindByOriginalLink(db *gorm.DB, linkOriginal string) (entities.ShortedLink, error) {
	shortedLink := entities.ShortedLink{}
	result := db.Where("link_original = ?", linkOriginal).First(&shortedLink)
	return shortedLink, result.Error
}
func (repository *ShortedLinkRepositoryImpl) FindByShortedLink(db *gorm.DB, generateLink string) (entities.ShortedLink, error) {
	shortedLink := entities.ShortedLink{}
	result := db.Where("link_shorted = ?", generateLink).First(&shortedLink)
	return shortedLink, result.Error
}
