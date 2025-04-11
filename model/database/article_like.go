package database

import "server/global"

type ArticleLike struct {
	global.MODEL
	ArticleID string `json:"article_id"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"-" gorm:"foreignKey:UserID"`
}
