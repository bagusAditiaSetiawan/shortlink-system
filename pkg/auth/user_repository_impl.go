package auth

import (
	"gorm.io/gorm"
	"shortlink-system/pkg/entities"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(tx *gorm.DB, user entities.User) (entities.User, error) {
	result := tx.Create(&user)
	return user, result.Error
}

func (r *UserRepositoryImpl) GetExisted(tx *gorm.DB, email string, username string) (entities.User, error) {
	var user entities.User
	result := tx.Where("email = ?", email).Or("username = ?", username).First(&user)
	return user, result.Error
}
