package httpsvc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidPayload              = echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
	ErrFailedToRegisterMailingList = echo.NewHTTPError(http.StatusInternalServerError, "failed to create mailing list")
)

func ErrCustomMsgAndStatus(status int, msg string) *echo.HTTPError {
	return echo.NewHTTPError(status, msg)
}
