package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Email string `gorm:"uniqure;not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"-" binding:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}