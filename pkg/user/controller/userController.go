package controller

import "github.com/labstack/echo/v4"

type UserContoller interface {
	GetByUserID(pctx echo.Context) error
	Update(pctx echo.Context) error
	Delete(pctx echo.Context) error
}
