package password

type Service interface {
	Hashing(password string) (string, error)
	CompareHashAndPassword(hash string, password string) error
}
