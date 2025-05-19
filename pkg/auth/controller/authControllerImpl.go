package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/jaayroots/todo-api/pkg/custom"
	"github.com/labstack/echo/v4"

	_authModel "github.com/jaayroots/todo-api/pkg/auth/model"
	_authService "github.com/jaayroots/todo-api/pkg/auth/service"
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
)

type authContollerImpl struct {
	authService _authService.AuthService
}

func NewAuthControllerImpl(
	authService _authService.AuthService,
) AuthContoller {
	return &authContollerImpl{
		authService: authService,
	}
}

func (c *authContollerImpl) Register(pctx echo.Context) error {

	createReq := new(_userModel.UserReq)

	customerEchoRequest := custom.NewCustomerEchoRequest(pctx)
	if err := customerEchoRequest.Build(createReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	user, err := c.authService.Register(createReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, user, "Register successful", nil)

}

func (c *authContollerImpl) Login(pctx echo.Context) error {

	loginReq := new(_authModel.LoginReq)

	customerEchoRequest := custom.NewCustomerEchoRequest(pctx)
	if err := customerEchoRequest.Build(loginReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	loginReq.IpAddress = pctx.RealIP()
	token, err := c.authService.Login(loginReq)
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, token, "Login successful", nil)

}

func (c *authContollerImpl) Logout(pctx echo.Context) error {

	val := pctx.Get("user")
	user, ok := val.(*_userModel.UserRes)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	err := c.authService.Logout(uint(user.ID))
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "Logout successful", nil)

}

func (c *authContollerImpl) Authorizing(pctx echo.Context, next echo.HandlerFunc) error {

	authHeader := pctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", nil)
	}

	token := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	}

	loginRes, err := c.authService.Authorizing(token)
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	pctx.Set("user", loginRes.User)

	ctx := context.WithValue(pctx.Request().Context(), "userID", loginRes.User.ID)
	req := pctx.Request().WithContext(ctx)

	pctx.SetRequest(req)

	return next(pctx)
}
