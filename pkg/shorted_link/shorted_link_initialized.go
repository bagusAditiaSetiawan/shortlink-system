package shorted_link

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"os"
	"shortlink-system/pkg/aws_cloudwatch"
	"shortlink-system/pkg/generator"
	"shortlink-system/pkg/redis"
	"strconv"
)

func InitializedShortedLinkService(db *gorm.DB, validate *validator.Validate, logger aws_cloudwatch.AwsCloudWatchService,
	redisService redis.RedisService,
	generatorService generator.Service) ShortLinkService {
	shortedLinkRepository := NewShortedLinkRepository()
	maxMonthlyUser := os.Getenv("MAX_MONTHLY_USER")
	maxMonthlyUserInt, _ := strconv.Atoi(maxMonthlyUser)
	return New(db, validate, redisService, shortedLinkRepository, generatorService, maxMonthlyUserInt, logger)
}
