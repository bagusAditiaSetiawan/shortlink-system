package jwt

import "github.com/golang-jwt/jwt/v5"

type JwtService interface {
	Verify(token string) (*TokenPayload, error)
	parse(token string) (*jwt.Token, error)
	Generate(payload *TokenPayload) string
}
