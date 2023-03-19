package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash []byte `json:"-"`
}

type Article struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
}
