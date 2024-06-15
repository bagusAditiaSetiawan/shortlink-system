package generator

import (
	"math/rand"
	"time"
)

type ShortService struct{}

func NewShortService() *ShortService {
	return &ShortService{}
}

func (s *ShortService) Generate() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[random.Intn(len(charset))]
	}
	return string(shortKey)
}
