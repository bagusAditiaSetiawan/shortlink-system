package auth

type Service interface {
	SignUp(req *SignUpRequest) SignUpSignature
	SignIn(req *SignInRequest) SignInResponse
}

