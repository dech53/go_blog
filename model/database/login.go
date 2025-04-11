package database

import "server/global"

type Login struct {
	global.MODEL
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	OS          string `json:"os"`
	DeviceInfo  string `json:"device_info"`
	BrowserInfo string `json:"browser_info"`
	Status      int    `json:"status"`
}
