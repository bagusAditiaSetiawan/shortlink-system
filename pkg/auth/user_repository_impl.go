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

func (r *UserRepositoryImpl) PaginateUser(db *gorm.DB, req *PaginateUser) ([]entities.User, int) {
	users := []entities.User{}
	var count int64
	query := db.Model(&entities.User{})
	if req.Username != "" {
		query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Email != "" {
		query.Where("email LIKE ?", "%"+req.Email+"%")
	}
	query.Count(&count)
	query.Offset(req.GetOffset()).Limit(req.Limit)
	query.Find(&users)
	return users, int(count)
}
