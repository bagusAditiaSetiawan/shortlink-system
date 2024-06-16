package auth

import "shortlink-system/api/presenter"

type PaginateUser struct {
	presenter.PaginatePresenter
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

type SignUpRequest struct {
	Username string `json:"username" form:"username" validate:"required,gte=3,lte=255"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,gte=8,lte=255"`
}

type SignUpSignature struct {
	Username string `json:"username" form:"username"`
}

type SignInRequest struct {
	Username string `json:"username" form:"username" validate:"required,gte=3,lte=255"`
	Password string `json:"password" form:"password" validate:"required,gte=8,lte=255"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}
