package models

type Article struct {
	BaseModel
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
}
