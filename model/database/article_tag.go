package database

type ArticleTag struct {
	Tag    string `json:"tag" gorm:"primaryKey"`
	Number int    `json:"number"`
}
