package httpsvc

import (
	"github.com/labstack/echo/v4"
)

type Service struct {
	group *echo.Group
}

func RouteService(group *echo.Group) {
	srv := &Service{
		group: group,
	}

	srv.initRoutes()
}

func (s *Service) initRoutes() {
	s.group.POST("/mail/free", s.registerFreeMail())
}
