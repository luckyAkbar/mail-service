package helper

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	sib "github.com/sendinblue/APIv3-go-library/lib"
	"github.com/sirupsen/logrus"
)

type SendInBlueHelper struct {
	api *sib.APIClient
}

var failedToSendEmail = echo.NewHTTPError(http.StatusInternalServerError, "Server failed to send email(s).")

func NewSIBHelper(sibCLient *sib.APIClient) *SendInBlueHelper {
	return &SendInBlueHelper{
		api: sibCLient,
	}
}

func (s *SendInBlueHelper) CreateEmailContent(
	subject,
	senderName,
	senderEmail,
	HtmlContent string,
	to sib.SendSmtpEmailTo,
) sib.SendSmtpEmail {
	return sib.SendSmtpEmail{
		HtmlContent: HtmlContent,
		Subject:     subject,
		Sender: &sib.SendSmtpEmailSender{
			Name:  senderName,
			Email: senderEmail,
		},
		To: []sib.SendSmtpEmailTo{
			to,
		},
	}
}

func (s *SendInBlueHelper) SendEmail(content sib.SendSmtpEmail) error {
	_, resp, err := s.api.TransactionalEmailsApi.SendTransacEmail(context.Background(), content)
	if resp.StatusCode != 201 { // error from sendinblue
		logrus.Error("Mailing utility broken, resp from Sendinblue:", resp)
		return failedToSendEmail
	}

	if err != nil {
		logrus.Error(err)
		return failedToSendEmail
	}

	logrus.Info(resp)

	return nil
}
