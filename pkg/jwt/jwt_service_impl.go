package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"shortlink-system/pkg/languages"
	"time"
)

type JwtServiceImpl struct {
	TokenKey        string
	TokenExpireTime string
}

func New(tokenKey string, tokenExpire string) *JwtServiceImpl {
	return &JwtServiceImpl{
		TokenKey:        tokenKey,
		TokenExpireTime: tokenExpire,
	}
}

type TokenPayload struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (service *JwtServiceImpl) Verify(token string) (*TokenPayload, error) {
	parsed, err := service.parse(token)

	if err != nil {
		return nil, err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New(languages.INTERNAL_ERROR)
	}
	return &TokenPayload{
		ID: uint(id),
	}, nil
}

func (service *JwtServiceImpl) parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(service.TokenKey), nil
	})
}
func (service *JwtServiceImpl) Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(service.TokenExpireTime)
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(v).Unix(),
		"id":       payload.ID,
		"username": payload.Username,
		"email":    payload.Email,
	})
	token, err := t.SignedString([]byte(service.TokenKey))

	if err != nil {
		panic(err)
	}

	return token
}
