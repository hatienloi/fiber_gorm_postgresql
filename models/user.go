package models

type User struct {
	BaseModel
	Name         string `json:"name" gorm:"not null;default;uniqueIndex"`
	Email        string `json:"email" gorm:"not null;default;uniqueIndex"`
	PasswordHash []byte `json:"-"`
}
