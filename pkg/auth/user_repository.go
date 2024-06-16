package auth

import (
	"gorm.io/gorm"
	"shortlink-system/pkg/entities"
)

type UserRepository interface {
	Create(tx *gorm.DB, user entities.User) (entities.User, error)
	GetExisted(tx *gorm.DB, email string, username string) (entities.User, error)
	PaginateUser(db *gorm.DB, req *PaginateUser) ([]entities.User, int)
}
