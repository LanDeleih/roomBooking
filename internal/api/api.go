package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type Api interface {
	Start() error
}

type Server struct {
	Server *echo.Echo
	Logger *logrus.Logger
}

func NewApi(logger *logrus.Logger) Api {
	n := echo.New()
	srv :=  &Server{
		Logger: logger,
		}
	srv.configureRoutes()

	return srv
}

func (s *Server) configureRoutes() {
	s.Server.Use(middleware.Recover())
	v1 := s.Server.Group("/api/v1")
	{
		v1.GET("create", func() error {
			return nil
		})

	}
}

func (s *Server) Start() error {
	return nil
}

