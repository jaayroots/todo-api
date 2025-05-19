package controller

import (
	"github.com/labstack/echo/v4"
)

type TodoContoller interface {
	Get(pctx echo.Context) error
	Create(pctx echo.Context) error
	Update(pctx echo.Context) error
	Delete(pctx echo.Context) error
	FindAll(pctx echo.Context) error
}
