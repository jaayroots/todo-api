package controller

import "github.com/labstack/echo/v4"

type AuthContoller interface {
	Register(pctx echo.Context) error
	Login(pctx echo.Context) error
	Logout(pctx echo.Context) error
	Authorizing(pctx echo.Context, next echo.HandlerFunc) error
}
