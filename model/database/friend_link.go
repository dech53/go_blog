package database

import "server/global"

type FriendLink struct {
	global.MODEL
	Logo        string `json:"logo" gorm:"size:255"`
	Image       Image  `json:"-" gorm:"foreignKey:Logo;references:URL"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
