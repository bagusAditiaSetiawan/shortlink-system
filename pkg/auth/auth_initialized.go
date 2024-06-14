package auth

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitializedAuthService(db *gorm.DB, validate *validator.Validate) Service {
	userRepository := NewUserRepositoryImpl()
	return NewServiceImpl(db, validate, userRepository)
}
