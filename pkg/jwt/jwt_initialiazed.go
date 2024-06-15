package jwt

import "os"

func InitializedJwt() *JwtServiceImpl {
	key := os.Getenv("TOKEN_KEY")
	expire := os.Getenv("TOKEN_EXPIRE")
	return New(key, expire)
}
