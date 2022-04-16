package repository

import (
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type freeMailRepository struct {
	db            *gorm.DB
	priorityLevel int
	status        int
	userID        int
	receipient    string
	mailerName    string
	payload       string
}

func NewFreeMailRepo(mailerName, receipient, payload string) *freeMailRepository {
	return &freeMailRepository{
		db:            db.DB,
		priorityLevel: config.LOWEST_PRIORITY_LEVEL,
		status:        config.MAILING_STATUS_PENDING,
		userID:        config.FREE_USER_ID,
		mailerName:    mailerName,
		receipient:    receipient,
		payload:       payload,
	}
}

func (f *freeMailRepository) Create() error {
	mailingList := &models.MailingList{
		Receipient: f.receipient,
		MailerName: f.mailerName,
		Payload:    f.payload,
		Priority:   f.priorityLevel,
		Status:     f.status,
		UserID:     f.userID,
	}

	if err := f.db.Create(mailingList).Error; err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
