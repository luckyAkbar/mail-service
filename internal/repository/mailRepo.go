package repository

import (
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MailRepository struct {
	db *gorm.DB
}

type MailInput struct {
	Subject         string
	ReceipientEmail string
	ReceipientName  string
	SenderName      string
	HTMLContent     string
	Priority        int
	Status          int
	UserID          int
}

func NewMailRepo() *MailRepository {
	return &MailRepository{
		db: db.DB,
	}
}

func (f *MailRepository) RegisterEmail(input MailInput) error {
	mailingList := &models.MailingList{
		Subject:         input.Subject,
		ReceipientEmail: input.ReceipientEmail,
		ReceipientName:  input.ReceipientName,
		SenderName:      input.SenderName,
		HtmlContent:     input.HTMLContent,
		Priority:        input.Priority,
		Status:          input.Status,
		UserID:          input.UserID,
	}

	if err := f.db.Create(mailingList).Error; err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (f *MailRepository) GetPendingFreeMailingList(limit int) ([]models.MailingList, error) {
	if limit == 0 {
		logrus.Warn("Query free mailing list called using limit = 0")
		limit = config.DEFAULT_FREE_EMAIL_LIST_QUERY_LIMIT
	}

	list := []models.MailingList{}

	err := f.db.Model(&models.MailingList{}).
		Where("priority = ? AND status = ? AND user_id = ?",
			config.LOWEST_PRIORITY_LEVEL,
			config.MAILING_STATUS_PENDING,
			config.FREE_USER_ID,
		).Limit(limit).
		Scan(&list).Error

	if err != nil {
		logrus.Error(err)
		return list, err
	}

	return list, nil
}

func (f *MailRepository) MarkAsSent(mail models.MailingList) {
	mail.Status = config.MAILING_STATUS_SUCCESS
	if err := f.db.Save(&mail).Error; err != nil {
		logrus.Error(err)
	}
}

func (f *MailRepository) MarkAsFailed(mail models.MailingList) {
	mail.Status = config.MAILING_STATUS_FAILED
	if err := f.db.Save(&mail).Error; err != nil {
		logrus.Error(err)
	}
}
