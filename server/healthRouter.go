package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *echoServer) healthRouter() {
	router := s.app.Group("/health")
	router.GET("", s.healthCheck)
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
