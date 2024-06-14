package entities

import "gorm.io/gorm"

type ShortedLink struct {
	gorm.Model
	UserID       int
	User         User
	LinkOriginal string `gorm:"varchar:256" json:"link_original"`
	LinkShorted  string `gorm:"varchar:256" json:"link_shorted"`
}
