package controller

import (
	"net/http"

	"github.com/jaayroots/todo-api/pkg/custom"
	"github.com/labstack/echo/v4"

	_userModel "github.com/jaayroots/todo-api/pkg/user/model"
	_userService "github.com/jaayroots/todo-api/pkg/user/service"
	_utils "github.com/jaayroots/todo-api/utils"
)

type userContollerImpl struct {
	userService _userService.UserService
}

func NewUserControllerImpl(
	userService _userService.UserService,
) UserContoller {
	return &userContollerImpl{
		userService,
	}
}

func (c *userContollerImpl) GetByUserID(pctx echo.Context) error {

	userId, err := _utils.StrToUint64(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	user, err := c.userService.GetByUserID(userId)
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, user, "", nil)

}

func (c *userContollerImpl) Update(pctx echo.Context) error {

	updateReq := new(_userModel.UserUpdateReq)

	customerEchoRequest := custom.NewCustomerEchoRequest(pctx)
	if err := customerEchoRequest.Bild(updateReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	userId, err := _utils.StrToUint64(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	err = c.userService.Update(userId, updateReq)
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "", nil)

}

func (c *userContollerImpl) Delete(pctx echo.Context) error {

	userId, err := _utils.StrToUint64(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	err = c.userService.Delete(userId)
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "", nil)

}
