package usecase

import (
	"fmt"
	"mail-service/internal/helper"
	"mail-service/internal/repository"

	"github.com/sirupsen/logrus"
)

type FreeMail struct {
	ReceipientName string
	MailerName     string
	Payload        string
}

func NewFreeMailHandler(receipientName, mailerName, payload string) *FreeMail {
	return &FreeMail{
		ReceipientName: receipientName,
		MailerName:     mailerName,
		Payload:        payload,
	}
}

func (m *FreeMail) Validate() error {
	err := helper.ValidateEmailAdressList([]string{
		m.ReceipientName,
		m.MailerName,
	})

	if err != nil {
		return fmt.Errorf("Operation invalid because: %s", err.Error())
	}

	return nil
}

func (m *FreeMail) Register() error {
	mailRepo := repository.NewFreeMailRepo(m.MailerName, m.ReceipientName, m.Payload)

	if err := mailRepo.Create(); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
