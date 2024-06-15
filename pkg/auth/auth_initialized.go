package auth

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"shortlink-system/pkg/password"
)

func InitializedAuthService(db *gorm.DB, validate *validator.Validate) Service {
	userRepository := NewUserRepositoryImpl()
	passwordService := password.NewPasswordService()
	return NewServiceImpl(db, validate, userRepository, passwordService)
}
