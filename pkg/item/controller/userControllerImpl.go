package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jaayroots/todo-api/pkg/custom"
	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"
	_itemService "github.com/jaayroots/todo-api/pkg/item/service"
	_userModel "github.com/jaayroots/todo-api/pkg/user/model"

	_utils "github.com/jaayroots/todo-api/utils"
)

type itemContollerImpl struct {
	itemService _itemService.ItemService
}

func NewItemControllerImpl(
	itemService _itemService.ItemService,
) ItemContoller {
	return &itemContollerImpl{
		itemService: itemService,
	}
}

func (c *itemContollerImpl) Create(pctx echo.Context) error {

	val := pctx.Get("user")
	_, ok := val.(*_userModel.UserRes)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid request")
	}

	itemReq := new(_itemModel.ItemReq)

	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(itemReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	ctx := pctx.Request().Context()
	item, err := c.itemService.Create(ctx, itemReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, item, "", nil)

}

func (c *itemContollerImpl) Get(pctx echo.Context) error {

	val := pctx.Get("user")
	_, ok := val.(*_userModel.UserRes)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid request")
	}

	itemId, err := _utils.StrToUint(pctx.Param("itemID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid itemId", nil)
	}

	ctx := pctx.Request().Context()
	todo, err := c.itemService.Get(ctx, itemId)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, todo, "", nil)

}
