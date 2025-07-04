package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"-"` // ignore this field in output
}