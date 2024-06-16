package entities

import "gorm.io/gorm"

type UserRole string

const (
	ADMIN UserRole = "ADMIN"
	USER  UserRole = "USER"
)

type User struct {
	gorm.Model
	Username string   `gorm:"varchar:256;unique"`
	Email    string   `gorm:"varchar:256;unique"`
	Password string   `gorm:"varchar:256"`
	Role     UserRole `gorm:"type:enum('ADMIN','USER'); default:'USER'"`
}
