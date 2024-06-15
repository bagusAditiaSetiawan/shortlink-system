package auth

import "shortlink-system/pkg/entities"

type Service interface {
	SignUp(req *SignUpRequest) SignUpSignature
	SignIn(req *SignInRequest) entities.User
}
