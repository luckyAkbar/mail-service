package usecase

import (
	"context"
	"mail-service/internal/model"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/kumparan/go-utils/encryption"
	"github.com/sirupsen/logrus"
)

type mailUsecase struct {
	mailRepo model.MailRepository
	cryptor  *encryption.AESCryptor
}

func NewMailUsecase(mailRepo model.MailRepository, cryptor *encryption.AESCryptor) model.MailUsecase {
	return &mailUsecase{
		mailRepo,
		cryptor,
	}
}

func (u *mailUsecase) RegisterFreeMailingList(ctx context.Context, input *model.RegisterFreeMailInput) (*model.MailingList, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         utils.DumpIncomingContext(ctx),
		"mailingList": utils.Dump(input),
	})

	if err := input.Validate(); err != nil {
		return nil, ErrValidation
	}

	mailingList := &model.MailingList{
		ID:              utils.GenerateID(),
		Subject:         input.Subject,
		ReceipientEmail: input.ReceipientEmail,
		ReceipientName:  input.ReceipientName,
		SenderEmail:     input.SenderEmail,
		SenderName:      input.SenderName,
		Content:         input.Content,
		CreatedAt:       time.Now(),
		Status:          model.StatusPending,
		Priority:        model.PriorityLow,
	}

	if err := mailingList.Encrypt(); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	if err := u.mailRepo.Create(ctx, mailingList); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	_ = mailingList.Decrypt()
	return mailingList, nil
}
