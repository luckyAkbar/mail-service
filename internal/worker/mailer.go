package worker

import (
	"context"
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/helper"
	"mail-service/internal/model"
	"time"

	sib "github.com/sendinblue/APIv3-go-library/lib"
	"gopkg.in/guregu/null.v4"

	"github.com/kumparan/go-utils/encryption"
	"github.com/sirupsen/logrus"
)

type MailWorker struct {
	mailRepo model.MailRepository
	sibApi   *helper.SendInBlueHelper
	Cryptor  *encryption.AESCryptor
}

func NewMailWorker(mailRepo model.MailRepository, sibApi *helper.SendInBlueHelper) *MailWorker {
	return &MailWorker{
		mailRepo: mailRepo,
		sibApi:   sibApi,
		Cryptor:  helper.CreateCryptor(),
	}
}

func (m *MailWorker) SpawnWorker() {
	logrus.Info("Mail Worker Spawned")
	for {
		time.Sleep(config.FreeMailerWorkerSleepDuration())
		logrus.Info("New Iteration for Mail Worker")

		list, err := m.mailRepo.GetPendingMailingList(context.Background(), config.MailProcessingLimit())
		if err != nil {
			logrus.Error("Mail Worker error:", err)
			continue
		}

		if len(list) == 0 {
			logrus.Println("No free mailing list were found on this iteration.")
			continue
		}

		m.sendEmail(list)
	}
}

func (m *MailWorker) sendEmail(list []model.MailingList) {
	for _, mail := range list {
		if err := mail.Decrypt(); err != nil {
			logrus.Error(err)
			continue
		}

		now := time.Now()
		err := m.sibApi.SendEmail(m.sibApi.CreateEmailContent(
			mail.Subject,
			mail.SenderName,
			mail.SenderEmail,
			mail.Content,
			sib.SendSmtpEmailTo{
				Name:  mail.ReceipientName,
				Email: mail.ReceipientEmail,
			},
		))

		mail.LastSendingAttempt = null.NewTime(now, true)

		if err := mail.Encrypt(); err != nil {
			logrus.Error("mail content encryption failed. storing unencrypted value on db", err.Error())
		}

		if err != nil {
			logrus.Error(err)
			m.mailRepo.MarkAsFailed(context.Background(), &mail)
			continue
		}

		mail.SentAt = null.NewTime(now, true)

		m.mailRepo.MarkAsSent(context.Background(), &mail)
	}

	logrus.Info(fmt.Sprintf("Processing %d free email in this iteration", len(list)))
}
