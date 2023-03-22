package models

type Article struct {
	BaseModel
	Title       string `gorm:"unique"`
	Description string
}
