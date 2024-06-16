package shorted_link

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"shortlink-system/pkg/entities"
	"time"
)

type ShortedLinkRepositoryImpl struct {
}

func NewShortedLinkRepository() *ShortedLinkRepositoryImpl {
	return &ShortedLinkRepositoryImpl{}
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

func (repository *ShortedLinkRepositoryImpl) SumMonthlyUser(db *gorm.DB, id int) int {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	lastOfMonth := nextMonth.Add(-time.Second)
	var count int64
	db.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		Model(&entities.ShortedLink{}).
		Where("user_id = ? and created_at >= ? and updated_at <= ?", id, firstOfMonth, lastOfMonth).
		Count(&count)
	return int(count)
}

func (repository *ShortedLinkRepositoryImpl) FindByShortedLinkWithLock(db *gorm.DB, link string) (entities.ShortedLink, error) {
	shortedLink := entities.ShortedLink{}
	result := db.Where("link_shorted = ?", link).First(&shortedLink)
	return shortedLink, result.Error
}

func (repository *ShortedLinkRepositoryImpl) Update(db *gorm.DB, req entities.ShortedLink) entities.ShortedLink {
	db.Save(&req)
	return req
}

func (repository *ShortedLinkRepositoryImpl) PaginateShortLink(db *gorm.DB, req *PaginateShortedLink) ([]entities.ShortedLink, int) {
	shortedLinks := make([]entities.ShortedLink, 0)
	var total int64
	query := db.Model(&entities.ShortedLink{})

	if len(req.Url) > 0 {
		query = query.Where("link_original LIKE ?", "%"+req.Url+"%")
	}
	query.Count(&total)
	query.Offset(req.GetOffset()).Limit(req.GetLimit())
	query.Find(&shortedLinks)

	return shortedLinks, int(total)
}
