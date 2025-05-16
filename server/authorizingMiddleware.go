package server

import (
	_authController "github.com/jaayroots/todo-api/pkg/auth/controller"
	"github.com/labstack/echo/v4"
)

type authorizingMiddleware struct {
	authController _authController.AuthContoller
}

func (m *authorizingMiddleware) Authorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.authController.Authorizing(pctx, next)
	}
}
