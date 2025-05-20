package controller

import "github.com/labstack/echo/v4"

type ItemContoller interface {
	Create(pctx echo.Context) error
	Get(pctx echo.Context) error
}
