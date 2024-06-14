package password

import "golang.org/x/crypto/bcrypt"

type PasswordServiceImpl struct{}

func (s *PasswordServiceImpl) Hashing(password string) (string, error) {
	hashedByte, err := bcrypt.Cost([]byte(password))
	return string(hashedByte), err
}

func (s *PasswordServiceImpl) CompareHashAndPassword(hash string, password string) error {
	compared := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return compared
}
