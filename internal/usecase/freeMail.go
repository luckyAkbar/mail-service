package usecase

import (
	"mail-service/internal/config"
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
	mailRepo := repository.NewMailRepo()
	input, err := m.composeEmailInput()
	if err != nil {
		logrus.Error(err)
		return err
	}

	if err := mailRepo.RegisterEmail(input); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (m *FreeMail) composeEmailInput() (repository.MailInput, error) {
	space := " "
	cryptor := helper.CreateCryptor()
	subject, err := cryptor.Encrypt(strings.Trim(m.Subject, space))
	if err != nil {
		logrus.Error(err)
		return repository.MailInput{}, err
	}

	receipientEmail, err := cryptor.Encrypt(strings.Trim(m.ReceipientEmail, space))
	if err != nil {
		logrus.Error(err)
		return repository.MailInput{}, err
	}

	receipientName, err := cryptor.Encrypt(strings.Trim(m.ReceipientName, space))
	if err != nil {
		logrus.Error(err)
		return repository.MailInput{}, err
	}

	senderName, err := cryptor.Encrypt(strings.Trim(m.SenderName, space))
	if err != nil {
		logrus.Error(err)
		return repository.MailInput{}, err
	}

	HTMLContent, err := cryptor.Encrypt(m.HTMLContent)
	if err != nil {
		logrus.Error(err)
		return repository.MailInput{}, err
	}

	return repository.MailInput{
		Subject:         subject,
		ReceipientEmail: receipientEmail,
		ReceipientName:  receipientName,
		SenderName:      senderName,
		HTMLContent:     HTMLContent,
		Priority:        config.LOWEST_PRIORITY_LEVEL,
		Status:          config.MAILING_STATUS_PENDING,
		UserID:          config.FREE_USER_ID,
	}, nil
}
