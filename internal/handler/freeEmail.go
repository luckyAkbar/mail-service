package handler

import (
	"mail-service/internal/config"
	"mail-service/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type singleMailPayload struct {
	Receipient string `json:"receipient"`
	MailerName string `json:"mailerName"`
	MailBody   string `json:"body"`
}

func RegisterSingleFreeMail(c echo.Context) error {
	payload := new(singleMailPayload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, FailedRegisterEmail{
			OK:      false,
			Message: "Invalid payload.",
		})
	}

	if payload.Receipient == "" || payload.MailerName == "" || payload.MailBody == "" {
		return c.JSON(http.StatusBadRequest, FailedRegisterEmail{
			OK:      false,
			Message: "All required field must not empty.",
		})
	}

	mailingListID, err := utils.RegisterSingleMailingList(
		payload.Receipient,
		payload.MailerName,
		payload.MailBody,
		config.LOWEST_PRIORITY_LEVEL,
		config.FREE_USER_ID,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, FailedRegisterEmail{
			OK:      false,
			Message: "Failed to register request to mailing list.",
		})
	}

	return c.JSON(http.StatusOK, SuccessRegisterFreeEmail{
		OK:            true,
		Receipient:    payload.Receipient,
		MailingListID: mailingListID,
	})
}
