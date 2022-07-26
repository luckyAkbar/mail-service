package repository

import (
	"context"
	"mail-service/internal/model"

	"github.com/kumparan/go-utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type mailRepository struct {
	db *gorm.DB
}

func NewMailRepository(db *gorm.DB) model.MailRepository {
	return &mailRepository{
		db,
	}
}

func (r *mailRepository) Create(ctx context.Context, input *model.MailingList) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         utils.DumpIncomingContext(ctx),
		"mailingList": utils.Dump(input),
	})

	err := r.db.WithContext(ctx).Create(input).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// GetPendingMailingList get pending mailing list and sort them based on it's priority
// high priority list will always returned first if according to limit
func (r *mailRepository) GetPendingMailingList(ctx context.Context, limit int) (lists []model.MailingList, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"limit": limit,
	})
	err = r.db.WithContext(ctx).Model(&model.MailingList{}).
		Where("status = ?", model.StatusPending).Order("priority ASC").Limit(limit).
		Scan(&lists).Error

	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *mailRepository) MarkAsSent(ctx context.Context, mail *model.MailingList) error {
	mail.Status = model.StatusSent
	if err := r.db.WithContext(ctx).Save(mail).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.DumpIncomingContext(ctx),
			"mail": utils.Dump(mail),
		}).Error(err)

		return err
	}

	return nil
}

func (r *mailRepository) MarkAsFailed(ctx context.Context, mail *model.MailingList) error {
	mail.Status = model.StatusFailed
	if err := r.db.WithContext(ctx).Save(mail).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.DumpIncomingContext(ctx),
			"mail": utils.Dump(mail),
		}).Error(err)

		return err
	}

	return nil
}
