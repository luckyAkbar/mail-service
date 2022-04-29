package usecase

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrSenderNameEmpty       = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: sender name is empty")
	ErrSubjectEmpty          = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: subject is empty")
	ErrHTMLContentEmpty      = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: HTML content is empty")
	ErrMailAdressInvalid     = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: email address is invalid")
	ErrMailAdressAlreadyUsed = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: email address is used")
	ErrPasswordMismatch      = echo.NewHTTPError(http.StatusBadRequest, "Password mismatch")
	ErrPasswordTooWeak       = echo.NewHTTPError(http.StatusBadRequest, "Password should atleast 8 chars long, with 1 uppercase and 1 special character")
)
