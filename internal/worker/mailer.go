package worker

import (
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/helper"
	"mail-service/internal/models"
	"mail-service/internal/repository"
	"time"

	sib "github.com/sendinblue/APIv3-go-library/lib"
	"github.com/sirupsen/logrus"
)

type MailWorker struct {
	MailRepo *repository.MailRepository
}

func NewMailWorker() *MailWorker {
	return &MailWorker{
		MailRepo: repository.NewMailRepo(),
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
		err := sibHelper.SendEmail(sibHelper.CreateEmailContent(
			mail.Subject,
			mail.SenderName,
			mail.HtmlContent,
			sib.SendSmtpEmailTo{
				Name:  mail.ReceipientName,
				Email: mail.ReceipientEmail,
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
