package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"not null; unique"`
	Username string `gorm:"not null; unique"`
}
