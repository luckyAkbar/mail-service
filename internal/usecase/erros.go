package usecase

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrSenderNameEmpty  = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: sender name is empty")
	ErrSubjectEmpty     = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: subject is empty")
	ErrHTMLContentEmpty = echo.NewHTTPError(http.StatusBadRequest, "Operation invalid because: HTML content is empty")
)
