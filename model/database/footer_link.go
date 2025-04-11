package database

type FooterLink struct {
	Title string `json:"title" gorm:"primaryKey"`
	Link  string `json:"link"`
}
