package auth

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/pkg/entities"
	"shortlink-system/pkg/helper"
	"shortlink-system/pkg/languages"
	"shortlink-system/pkg/password"
)

type ServiceImpl struct {
	DB              *gorm.DB
	Validator       *validator.Validate
	UserRepository  UserRepository
	PasswordService password.Service
}

func NewServiceImpl(db *gorm.DB, validate *validator.Validate, userRepository UserRepository) *ServiceImpl {
	return &ServiceImpl{
		DB:             db,
		UserRepository: userRepository,
		Validator:      validate,
	}
}

func (service *ServiceImpl) SignUp(req *SignUpRequest) SignUpSignature {
	tx := service.DB.Begin()
	defer helper.RollbackOrCommitDb(tx)

	_, err := service.UserRepository.GetExisted(tx, req.Email, req.Username)
	if err == nil {
		panic(exception.NewBadRequestException(languages.USER_EXIST))
	}
	hashedPassword, err := service.PasswordService.Hashing(req.Password)
	if err != nil {
		panic(languages.INTERNAL_ERROR)
	}
	user := entities.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	user, err = service.UserRepository.Create(tx, user)

	helper.IfErrorHandler(err)

	return SignUpSignature{
		Username: user.Username,
	}
}

func (service *ServiceImpl) SignIn(req *SignInRequest) SignInResponse {
	panic("implement me")
}
