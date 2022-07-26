package model

import (
	"context"
	"mail-service/internal/helper"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

type Status string

var (
	StatusPending Status = "PENDING"
	StatusSent    Status = "SENT"
	StatusFailed  Status = "FAILED"
)

type Priority string

var (
	PriorityLow  Priority = "LOW"
	PriorityHigh Priority = "HIGH"
)

type MailingList struct {
	ID                 int64     `json:"id" gorm:"primaryKey"`
	Subject            string    `json:"subject" gorm:"not null"`
	ReceipientEmail    string    `json:"receipient_email" gorm:"not null"`
	ReceipientName     string    `json:"receipient_name" gorm:"not null"`
	SenderEmail        string    `json:"sender_email" gorm:"not null"`
	SenderName         string    `json:"sender_name" gorm:"not null"`
	Content            string    `json:"content" gorm:"not null"`
	CreatedAt          time.Time `json:"created_at" gorm:"not null"`
	Status             Status    `json:"status" gorm:"not null"`
	Priority           Priority  `json:"priority" gorm:"not null"`
	SentAt             null.Time `json:"sent_at"`
	LastSendingAttempt null.Time `json:"last_sending_attempt"`
}

func (m *MailingList) Encrypt() error {
	cryptor := helper.CreateCryptor()
	subject, err := cryptor.Encrypt(m.Subject)
	if err != nil {
		logrus.Error("encryption failed:", err.Error())
		return err
	}
	m.Subject = subject

	receipientEmail, err := cryptor.Encrypt(m.ReceipientEmail)
	if err != nil {
		logrus.Error("encryption failed:", err.Error())
		return err
	}
	m.ReceipientEmail = receipientEmail

	receipientName, err := cryptor.Encrypt(m.ReceipientName)
	if err != nil {
		logrus.Error("encryption failed:", err.Error())
		return err
	}
	m.ReceipientName = receipientName

	senderEmail, err := cryptor.Encrypt(m.SenderEmail)
	if err != nil {
		logrus.Error("encryption failed:", err.Error())
		return err
	}
	m.SenderEmail = senderEmail

	senderName, err := cryptor.Encrypt(m.SenderName)
	if err != nil {
		logrus.Error("encryption failed:", err.Error())
		return err
	}
	m.SenderName = senderName

	content, err := cryptor.Encrypt(m.Content)
	if err != nil {
		logrus.Error("encryption failed:", err.Error())
		return err
	}
	m.Content = content

	return nil
}

func (m *MailingList) Decrypt() error {
	cryptor := helper.CreateCryptor()
	subject, err := cryptor.Decrypt(m.Subject)
	if err != nil {
		logrus.Error("decryption failed:", err.Error())
		return err
	}
	m.Subject = subject

	receipientEmail, err := cryptor.Decrypt(m.ReceipientEmail)
	if err != nil {
		logrus.Error("decryption failed:", err.Error())
		return err
	}
	m.ReceipientEmail = receipientEmail

	receipientName, err := cryptor.Decrypt(m.ReceipientName)
	if err != nil {
		logrus.Error("decryption failed:", err.Error())
		return err
	}
	m.ReceipientName = receipientName

	senderEmail, err := cryptor.Decrypt(m.SenderEmail)
	if err != nil {
		logrus.Error("decryption failed:", err.Error())
		return err
	}
	m.SenderEmail = senderEmail

	senderName, err := cryptor.Decrypt(m.SenderName)
	if err != nil {
		logrus.Error("decryption failed:", err.Error())
		return err
	}
	m.SenderName = senderName

	content, err := cryptor.Decrypt(m.Content)
	if err != nil {
		logrus.Error("decryption failed:", err.Error())
		return err
	}
	m.Content = content

	return nil
}

type RegisterFreeMailInput struct {
	ReceipientEmail string `json:"receipient_email" validate:"required,email"`
	ReceipientName  string `json:"receipient_name" validate:"required"`
	SenderEmail     string `json:"sender_email" validate:"required,email"`
	SenderName      string `json:"sender_name" validate:"required"`
	Subject         string `json:"subject" validate:"required"`
	Content         string `json:"content" validate:"required"`
}

func (i *RegisterFreeMailInput) Validate() error {
	return validate.Struct(i)
}

type MailUsecase interface {
	RegisterFreeMailingList(ctx context.Context, input *RegisterFreeMailInput) (*MailingList, error)
}

type MailRepository interface {
	Create(ctx context.Context, mailingList *MailingList) error
	GetPendingMailingList(ctx context.Context, limit int) ([]MailingList, error)
	MarkAsSent(ctx context.Context, mailingList *MailingList) error
	MarkAsFailed(ctx context.Context, mailingList *MailingList) error
}
