package password

import "golang.org/x/crypto/bcrypt"

type PasswordServiceImpl struct{}

func NewPasswordService() *PasswordServiceImpl {
	return &PasswordServiceImpl{}
}

func (s *PasswordServiceImpl) Hashing(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedByte), err
}

func (s *PasswordServiceImpl) CompareHashAndPassword(hash string, password string) error {
	compared := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return compared
}
