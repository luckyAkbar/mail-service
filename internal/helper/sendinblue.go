package helper

import (
	"context"
	"mail-service/internal/config"
	"net/http"

	"github.com/labstack/echo/v4"
	sib "github.com/sendinblue/APIv3-go-library/lib"
	"github.com/sirupsen/logrus"
)

type SendInBlueHelper struct{}

var failedToSendEmail = echo.NewHTTPError(http.StatusInternalServerError, "Server failed to send email(s).")

func NewSIBHelper() *SendInBlueHelper {
	return &SendInBlueHelper{}
}

func (s *SendInBlueHelper) CreateEmailContent(
	subject,
	senderName,
	HtmlContent string,
	to sib.SendSmtpEmailTo,
) sib.SendSmtpEmail {
	return sib.SendSmtpEmail{
		HtmlContent: HtmlContent,
		Subject:     subject,
		Sender: &sib.SendSmtpEmailSender{
			Name:  senderName,
			Email: config.SIBRegisteredEmail(),
		},
		To: []sib.SendSmtpEmailTo{
			to,
		},
	}
}

func (s *SendInBlueHelper) SendEmail(content sib.SendSmtpEmail) error {
	var ctx context.Context
	sibClient := config.SIBClient()

	_, resp, err := sibClient.TransactionalEmailsApi.SendTransacEmail(ctx, content)
	if resp.StatusCode != 201 { // error from sendinblue
		logrus.Error("Mailing utility broken, resp from Sendinblue:", resp)
		return failedToSendEmail
	}

	if err != nil {
		logrus.Error(err)
		return failedToSendEmail
	}

	return nil
}
