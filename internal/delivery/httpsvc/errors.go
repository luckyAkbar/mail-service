package httpsvc

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidPayload = echo.NewHTTPError(http.StatusBadRequest, "Payload Invalid")
	ErrValidation     = echo.NewHTTPError(http.StatusBadRequest, "Validation Error")
	ErrInternal       = echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
)
