package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type User struct {
	BaseModel
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash []byte `json:"-"`
}

type Article struct {
	BaseModel
	Title       string `gorm:"unique"`
	Description string
}
