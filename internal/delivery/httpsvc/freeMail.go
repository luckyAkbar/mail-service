package httpsvc

import (
	"mail-service/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *Service) registerFreeMail() echo.HandlerFunc {
	type request struct {
		ReceipientName  string `json:"receipient_name"`
		ReceipientEmail string `json:"receipient_email"`
		SenderName      string `json:"sender_name"`
		HTMLContent     string `json:"html_content"`
		Subject         string `json:"subject"`
	}

	return func(c echo.Context) error {
		req := request{}

		if err := c.Bind(&req); err != nil {
			logrus.Error(err)
			return ErrInvalidPayload
		}

		handler := usecase.NewFreeMailHandler(
			req.ReceipientName,
			req.ReceipientEmail,
			req.SenderName,
			req.HTMLContent,
			req.Subject,
		)

		if err := handler.Validate(); err != nil {
			logrus.Error(err)
			return err
		}

		if err := handler.Register(); err != nil {
			logrus.Error(err)
			return ErrFailedToRegisterMailingList
		}

		return c.JSON(http.StatusOK, req)
	}
}
