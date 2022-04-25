package usecase

import (
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
	err := helper.ValidateEmailAdressList([]string{m.ReceipientEmail})
	if err != nil {
		return ErrMailAdressInvalid
	}

	if m.SenderName == "" {
		return ErrSenderNameEmpty
	}

	if m.Subject == "" {
		return ErrSubjectEmpty
	}

	if m.HTMLContent == "" {
		return ErrHTMLContentEmpty
	}

	return nil
}

func (m *FreeMail) Register() error {
	space := " "
	mailRepo := repository.NewMailRepo()
	err := mailRepo.RegisterFreeEmail(
		strings.Trim(m.Subject, space),
		strings.Trim(m.ReceipientEmail, space),
		strings.Trim(m.ReceipientName, space),
		strings.Trim(m.SenderName, space),
		m.HTMLContent,
	)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
