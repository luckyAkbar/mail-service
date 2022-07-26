package httpsvc

import (
	"mail-service/internal/model"

	"github.com/labstack/echo/v4"
)

type Service struct {
	mailUsecase model.MailUsecase
	group       *echo.Group
}

func InitService(group *echo.Group, mailUsecase model.MailUsecase) {
	srv := &Service{
		mailUsecase,
		group,
	}

	srv.initRoutes()
}

func (s *Service) initRoutes() {
	s.group.POST("/mail/free/", s.registerFreeMail())
}
