package httpsvc

import (
	"mail-service/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *Service) registerFreeMail() echo.HandlerFunc {
	type request struct {
		ReceipientName string `json:"receipient_name"`
		MailerName     string `json:"mailer_name"`
		Payload        string `json:"payload"`
	}

	return func(c echo.Context) error {
		req := request{}

		if err := c.Bind(&req); err != nil {
			logrus.Error(err)
			return ErrInvalidPayload
		}

		handler := usecase.NewFreeMailHandler(
			req.ReceipientName,
			req.MailerName, req.Payload,
		)

		if err := handler.Validate(); err != nil {
			logrus.Error(err)
			return ErrCustomMsgAndStatus(http.StatusBadRequest, err.Error())
		}

		if err := handler.Register(); err != nil {
			logrus.Error(err)
			return ErrFailedToRegisterMailingList
		}

		return c.JSON(http.StatusOK, req)
	}
}
