package entities

import "gorm.io/gorm"

type ShortedLink struct {
	gorm.Model
	UserID          uint
	User            User
	LinkOriginal    string `gorm:"varchar:256" json:"link_original"`
	LinkShorted     string `gorm:"varchar:256" json:"link_shorted"`
	LinkShortedFull string `gorm:"varchar:256" json:"link_shorted"`
	NumOfAccessed   int    `gorm:"int" json:"num_of_accessed"`
}
