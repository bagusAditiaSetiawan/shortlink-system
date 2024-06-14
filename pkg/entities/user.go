package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"varchar:256;unique"`
	Email    string `gorm:"varchar:256;unique"`
	Password string `gorm:"varchar:256"`
}
