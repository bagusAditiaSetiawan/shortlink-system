package auth

type SignUpRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignUpSignature struct {
	Username string `json:"username" form:"username"`
}

type SignInRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}
