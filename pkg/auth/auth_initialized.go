package auth

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"shortlink-system/pkg/aws_cloudwatch"
	"shortlink-system/pkg/password"
)

func InitializedAuthService(db *gorm.DB, validate *validator.Validate, logger aws_cloudwatch.AwsCloudWatchService) Service {
	userRepository := NewUserRepositoryImpl()
	passwordService := password.NewPasswordService()
	return NewServiceImpl(db, validate, userRepository, passwordService, logger)
}
