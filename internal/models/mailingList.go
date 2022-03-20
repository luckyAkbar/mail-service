package models

import "gorm.io/gorm"

type MailingList struct {
	gorm.Model

	Receipient string `json:"receipient" gorm:"not null"`
	MailerName string `json:"mailerName" gorm:"not null"`
	Payload    string `json:"payload" gorm:"not null"`
	Priority   int    `json:"priority" gorm:"default: 0"`
	Status     int    `json:"status" gorm:"default: 1"` // default is pending
	UserID     int

	User User `gorm:"foreignKey:UserID"`
}
