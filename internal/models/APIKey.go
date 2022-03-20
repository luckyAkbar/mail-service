package models

import "gorm.io/gorm"

type APIKey struct {
	gorm.Model

	Key    string `gorm:"not null"`
	UserID int

	User User `gorm:"foreignKey:UserID"`
}
