package custom

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseFormat struct {
	Status  string      `json:"status"`            // success หรือ unsuccess
	Code    int         `json:"code"`              // HTTP status code
	Message string      `json:"message,omitempty"` // ข้อความอธิบาย (optional)
	Data    interface{} `json:"data,omitempty"`    // ข้อมูล response (optional)
	Errors  string      `json:"errors,omitempty"`  // ข้อความ error (optional)
}

func Response(pctx echo.Context, httpStatus int, data interface{}, message string, err error) error {
	statusCode := "success"
	if httpStatus != http.StatusOK {
		statusCode = "unsuccess"
	}

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	res := ResponseFormat{
		Status:  statusCode,
		Code:    httpStatus,
		Message: message,
		Data:    data,
		Errors:  errMsg,
	}

	return pctx.JSON(httpStatus, res)
}
