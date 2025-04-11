package database

import (
	"server/global"
	"server/model/appTypes"
)

type Image struct {
	global.MODEL
	Name     string            `json:"name"`
	URL      string            `json:"url" gorm:"size:255;unique"`
	Category appTypes.Category `json:"category"`
	Storage  appTypes.Storage  `json:"storage"`
}
