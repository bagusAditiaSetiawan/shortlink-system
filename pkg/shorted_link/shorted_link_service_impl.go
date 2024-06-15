package shorted_link

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"shortlink-system/pkg/entities"
	"shortlink-system/pkg/generator"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/redis"
	"time"
)

type ShortLinkServiceImpl struct {
	DB                  *gorm.DB
	Validate            *validator.Validate
	RedisService        redis.RedisService
	ShortLinkRepository ShortedLinkRepository
	ShortLinkGenerator  generator.Service
}

func New(db *gorm.DB, validate *validator.Validate, redis redis.RedisService, repository ShortedLinkRepository, service generator.Service) *ShortLinkServiceImpl {
	return &ShortLinkServiceImpl{
		DB:                  db,
		Validate:            validate,
		RedisService:        redis,
		ShortLinkRepository: repository,
		ShortLinkGenerator:  service,
	}
}

func (service *ShortLinkServiceImpl) Handler(req *CreateShortedLink, user entities.User) entities.ShortedLink {
	err := service.Validate.Struct(req)
	helper.IfErrorHandler(err)
	shortedLink, err := service.GetExistShortLink(req.Url)
	if err != nil {
		shortedLink = service.ShortLinkGenerator.Generate()
	}
	tx := service.DB.Begin()
	dataShortedLink := service.ShortLinkRepository.Create(tx, entities.ShortedLink{
		LinkShorted:  shortedLink,
		LinkOriginal: req.Url,
		UserID:       user.ID,
	})
	go service.SaveToRedis(dataShortedLink)
	defer helper.RollbackOrCommitDb(tx)
	return dataShortedLink
}

func (service *ShortLinkServiceImpl) SaveToRedis(link entities.ShortedLink) {
	service.RedisService.Set(link.LinkOriginal, link.LinkShorted, 24*time.Hour)
	service.RedisService.Set(link.LinkShorted, link.LinkOriginal, 24*time.Hour)
}

func (service *ShortLinkServiceImpl) GetExistShortLink(linkOriginal string) (string, error) {
	if redisLink, err := service.RedisService.Get(linkOriginal); err == nil {
		if shortedLink, ok := redisLink.(string); ok {
			return shortedLink, nil
		}
	}
	shortedLink, err := service.ShortLinkRepository.FindByOriginalLink(service.DB, linkOriginal)
	if err != nil {
		return "", err
	}
	return shortedLink.LinkShorted, nil
}
func (service *ShortLinkServiceImpl) GetExistOriginalLink(generatedLink string) (string, error) {
	if redisLink, err := service.RedisService.Get(generatedLink); err == nil {
		if shortedLink, ok := redisLink.(string); ok {
			return shortedLink, nil
		}
	}
	shortedLink, err := service.ShortLinkRepository.FindByShortedLink(service.DB, generatedLink)
	if err != nil {
		return "", err
	}
	return shortedLink.LinkShorted, nil
}
