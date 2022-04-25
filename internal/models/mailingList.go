package models

import "gorm.io/gorm"

type MailingList struct {
	gorm.Model

	Subject         string `json:"subject" gorm:"not null"`
	ReceipientEmail string `json:"receipient_email" gorm:"not null"`
	ReceipientName  string `json:"receipient_name" gorm:"not null"`
	SenderName      string `json:"mailer_name" gorm:"not null"`
	HtmlContent     string `json:"html_content" gorm:"not null"`
	Priority        int    `json:"priority" gorm:"default: 0"`
	Status          int    `json:"status" gorm:"default: 1"` // default is pending
	UserID          int

	User User `gorm:"foreignKey:UserID"`
}
