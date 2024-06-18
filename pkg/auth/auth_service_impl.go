package auth

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	exception "shortlink-system/api/exceptions"
	"shortlink-system/pkg/aws_cloudwatch"
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
	Logger          aws_cloudwatch.AwsCloudWatchService
}

func NewServiceImpl(db *gorm.DB, validate *validator.Validate, userRepository UserRepository, passwordService password.Service, logger aws_cloudwatch.AwsCloudWatchService) *ServiceImpl {
	return &ServiceImpl{
		DB:              db,
		UserRepository:  userRepository,
		Validator:       validate,
		PasswordService: passwordService,
		Logger:          logger,
	}
}

func (service *ServiceImpl) SignUp(req *SignUpRequest) SignUpSignature {
	err := service.Validator.Struct(req)
	helper.IfErrorHandler(err)
	tx := service.DB.Begin()
	defer helper.RollbackOrCommitDb(tx)
	go service.Logger.SendLogInfo("Get existed user")
	_, err = service.UserRepository.GetExisted(tx, req.Email, req.Username)
	go service.Logger.SendLogInfo("Result get existed", err)

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
	go service.Logger.SendLogInfo("Create user repository")
	user, err = service.UserRepository.Create(tx, user)

	go service.Logger.SendLogInfo("Result error create", err)

	helper.IfErrorHandler(err)
	return SignUpSignature{
		Username: user.Username,
	}
}

func (service *ServiceImpl) SignIn(req *SignInRequest) entities.User {
	err := service.Validator.Struct(req)
	helper.IfErrorHandler(err)
	tx := service.DB.Begin()
	defer helper.RollbackOrCommitDb(tx)
	user, err := service.UserRepository.GetExisted(tx, "", req.Username)
	if err != nil {
		panic(exception.NewNotFoundHandler(languages.USER_NOT_FOUND))
	}
	if matched := service.PasswordService.CompareHashAndPassword(user.Password, req.Password); matched != nil {
		panic(exception.NewBadRequestException(languages.PASSWORD_WRONG))
	}
	return user
}
