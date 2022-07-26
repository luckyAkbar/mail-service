package httpsvc

import (
	"mail-service/internal/model"
	"mail-service/internal/usecase"
	"net/http"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *Service) registerFreeMail() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &model.RegisterFreeMailInput{}

		if err := c.Bind(req); err != nil {
			logrus.Error(err)
			return ErrInvalidPayload
		}

		mailingList, err := s.mailUsecase.RegisterFreeMailingList(c.Request().Context(), req)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(c.Request().Context()),
				"input": utils.Dump(req),
			}).Error(err)

			return ErrInternal
		case usecase.ErrValidation:
			return ErrValidation
		case nil:
			return c.JSON(http.StatusOK, mailingList)
		}
	}
}
