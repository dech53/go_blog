package database

import (
	"github.com/gofrs/uuid"
	"server/global"
	"server/model/appTypes"
)

type User struct {
	global.MODEL
	UUID      uuid.UUID       `json:"uuid" gorm:"type:char(36);unique"`
	Username  string          `json:"username"`
	Password  string          `json:"-"`
	Email     string          `json:"email"`
	Avatar    string          `json:"avatar" gorm:"size:255"`
	Address   string          `json:"address"`
	Signature string          `json:"signature" gorm:"default:'签名是空白的，这位用户似乎比较低调。'"`
	RoleId    appTypes.RoleID `json:"role_id"`
	Freeze    bool            `json:"freeze"`
}
