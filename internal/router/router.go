package router

import (
	"mail-service/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", handler.Home)
	e.POST("/free/sendEmail/single", handler.RegisterSingleFreeMail)

	return e
}
