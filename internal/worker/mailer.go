package worker

import (
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/helper"
	"mail-service/internal/models"
	"mail-service/internal/repository"
	"time"

	"github.com/kumparan/go-utils/encryption"
	sib "github.com/sendinblue/APIv3-go-library/lib"
	"github.com/sirupsen/logrus"
)

type MailWorker struct {
	MailRepo *repository.MailRepository
	Cryptor  *encryption.AESCryptor
}

func NewMailWorker() *MailWorker {
	return &MailWorker{
		MailRepo: repository.NewMailRepo(),
		Cryptor:  helper.CreateCryptor(),
	}
}

func (m *MailWorker) SpawnWorker() {
	logrus.Info("Mail Worker Spawned")
	for true {
		time.Sleep(config.FreeMailerWorkerSleepDuration())
		logrus.Info("New Iteration for Mail Worker")

		list, err := m.MailRepo.GetPendingFreeMailingList(config.FreeMailProcessingLimit())
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

func (m *MailWorker) sendEmail(list []models.MailingList) {
	sibHelper := helper.NewSIBHelper()
	for _, mail := range list {
		content, err := m.decryptEmailContent(mail)
		if err != nil {
			logrus.Error(err)
			continue
		}

		err = sibHelper.SendEmail(sibHelper.CreateEmailContent(
			content.Subject,
			content.SenderName,
			content.HtmlContent,
			sib.SendSmtpEmailTo{
				Name:  content.ReceipientName,
				Email: content.ReceipientEmail,
			},
		))

		if err != nil {
			logrus.Error(err)
			m.MailRepo.MarkAsFailed(mail)
			continue
		}

		m.MailRepo.MarkAsSent(mail)
	}

	logrus.Info(fmt.Sprintf("Processing %d free email in this iteration", len(list)))
}

func (m *MailWorker) decryptEmailContent(content models.MailingList) (models.MailingList, error) {
	subject, err := m.Cryptor.Decrypt(content.Subject)
	if err != nil {
		logrus.Error(err)
		return content, err
	}

	receipientEmail, err := m.Cryptor.Decrypt(content.ReceipientEmail)
	if err != nil {
		logrus.Error(err)
		return content, err
	}

	receipientName, err := m.Cryptor.Decrypt(content.ReceipientName)
	if err != nil {
		logrus.Error(err)
		return content, err
	}

	senderName, err := m.Cryptor.Decrypt(content.SenderName)
	if err != nil {
		logrus.Error(err)
		return content, err
	}

	HTMLContent, err := m.Cryptor.Decrypt(content.HtmlContent)
	if err != nil {
		logrus.Error(err)
		return content, err
	}

	return models.MailingList{
		Subject:         subject,
		ReceipientEmail: receipientEmail,
		ReceipientName:  receipientName,
		SenderName:      senderName,
		HtmlContent:     HTMLContent,
	}, nil
}
