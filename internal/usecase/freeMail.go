package usecase

import (
	"fmt"
	"mail-service/internal/helper"
	"mail-service/internal/repository"
	"strings"

	"github.com/sirupsen/logrus"
)

type FreeMail struct {
	ReceipientName  string
	ReceipientEmail string
	SenderName      string
	HTMLContent     string
	Subject         string
}

func NewFreeMailHandler(
	receipientName,
	receipientEmail,
	senderName,
	HTMLContent,
	subject string,
) *FreeMail {
	return &FreeMail{
		ReceipientName:  receipientName,
		ReceipientEmail: receipientEmail,
		SenderName:      senderName,
		HTMLContent:     HTMLContent,
		Subject:         subject,
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
