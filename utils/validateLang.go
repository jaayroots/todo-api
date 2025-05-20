package utils

import (
	"github.com/labstack/echo/v4"
)

func ValidateLangOrDefault(pctx echo.Context) string {
	lang := pctx.Param("lang")
	if lang != "en" && lang != "th" {
		return "en"
	}
	return lang
}
