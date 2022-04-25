package repository

import (
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MailRepository struct {
	db            *gorm.DB
	priorityLevel int
	status        int
	userID        int
	receipient    string
	mailerName    string
	payload       string
}

func NewMailRepo() *MailRepository {
	return &MailRepository{
		db: db.DB,
	}
}

func (f *MailRepository) RegisterFreeEmail(subject, receipientEmail, receipientName, senderName, HTMLContent string) error {
	mailingList := &models.MailingList{
		Subject:         subject,
		ReceipientEmail: receipientEmail,
		ReceipientName:  receipientName,
		SenderName:      senderName,
		HtmlContent:     HTMLContent,
		Priority:        config.LOWEST_PRIORITY_LEVEL,
		Status:          config.MAILING_STATUS_PENDING,
		UserID:          config.FREE_USER_ID,
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
