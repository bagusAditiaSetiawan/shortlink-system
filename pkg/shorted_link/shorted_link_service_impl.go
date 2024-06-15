package shorted_link

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"os"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/pkg/auth"
	"shortlink-system/pkg/entities"
	"shortlink-system/pkg/generator"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/languages"
	"shortlink-system/pkg/redis"
	"strconv"
	"time"
)

type ShortLinkServiceImpl struct {
	DB                  *gorm.DB
	Validate            *validator.Validate
	RedisService        redis.RedisService
	ShortLinkRepository ShortedLinkRepository
	ShortLinkGenerator  generator.Service
	MaxMonthly          int
}

func New(db *gorm.DB, validate *validator.Validate, redis redis.RedisService, repository ShortedLinkRepository, service generator.Service, max int) *ShortLinkServiceImpl {
	return &ShortLinkServiceImpl{
		DB:                  db,
		Validate:            validate,
		RedisService:        redis,
		ShortLinkRepository: repository,
		ShortLinkGenerator:  service,
		MaxMonthly:          max,
	}
}

func (service *ShortLinkServiceImpl) CreateShortedLink(req *CreateShortedLink, user auth.UserLoggedPayload) entities.ShortedLink {
	err := service.Validate.Struct(req)
	helper.IfErrorHandler(err)
	tx := service.DB.Begin()
	defer helper.RollbackOrCommitDb(tx)

	SumLinkOfUser := service.ShortLinkRepository.SumMonthlyUser(tx, int(user.ID))
	if SumLinkOfUser > service.MaxMonthly-1 {
		panic(exception.NewBadRequestException(languages.MAX_USER_MONTHLY + "_" + strconv.Itoa(service.MaxMonthly)))
	}
	shortedLink := service.GenerateShortLink(tx)
	dataShortedLink := service.ShortLinkRepository.Create(tx, entities.ShortedLink{
		LinkShortedFull: fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), shortedLink),
		LinkShorted:     shortedLink,
		LinkOriginal:    req.Url,
		UserID:          user.ID,
	})
	go service.SaveToRedis(dataShortedLink)
	return dataShortedLink
}

func (service *ShortLinkServiceImpl) GenerateShortLink(tx *gorm.DB) string {
	shortedLink := service.ShortLinkGenerator.Generate()
	_, err := service.ShortLinkRepository.FindByShortedLink(tx, shortedLink)
	if err == nil {
		return service.GenerateShortLink(tx)
	}
	return shortedLink
}

func (service *ShortLinkServiceImpl) SaveToRedis(link entities.ShortedLink) {
	service.RedisService.Set(link.LinkShorted, link.LinkOriginal, 24*time.Hour)
}
func (service *ShortLinkServiceImpl) GetExistOriginalLink(generatedLink string) (string, error) {
	if redisLink, err := service.RedisService.Get(generatedLink); err == nil {
		if originalLink, ok := redisLink.(string); ok {
			log.Info("Get link from redis")
			return originalLink, nil
		}
	}
	log.Info("Get link from DB")
	shortedLink, err := service.ShortLinkRepository.FindByShortedLink(service.DB, generatedLink)
	return shortedLink.LinkOriginal, err
}

func (service *ShortLinkServiceImpl) UpdateAccessed(link string) entities.ShortedLink {
	tx := service.DB.Begin()
	defer helper.RollbackOrCommitDb(tx)
	shortedLink, err := service.ShortLinkRepository.FindByShortedLinkWithLock(tx, link)
	if err != nil {
		return shortedLink
	}
	shortedLink.NumOfAccessed = shortedLink.NumOfAccessed + 1
	service.ShortLinkRepository.Update(tx, shortedLink)
	return shortedLink
}
