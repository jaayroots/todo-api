package controller

import (
	"net/http"

	"github.com/jaayroots/todo-api/pkg/custom"
	"github.com/labstack/echo/v4"

	_todoModel "github.com/jaayroots/todo-api/pkg/todo/model"
	_todoService "github.com/jaayroots/todo-api/pkg/todo/service"
	_utils "github.com/jaayroots/todo-api/utils"
)

type todoContollerImpl struct {
	todoService _todoService.TodoService
}

func NewTodoControllerImpl(
	todoService _todoService.TodoService,
) TodoContoller {
	return &todoContollerImpl{
		todoService,
	}
}

func (c *todoContollerImpl) Get(pctx echo.Context) error {

	todoID, err := _utils.StrToUint(pctx.Param("todoID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid todoID", nil)
	}

	ctx := pctx.Request().Context()
	todo, err := c.todoService.Get(ctx, todoID)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, todo, "", nil)

}

func (c *todoContollerImpl) Create(pctx echo.Context) error {

	todoReq := new(_todoModel.TodoReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(todoReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	ctx := pctx.Request().Context()
	_, err := c.todoService.Create(ctx, todoReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *todoContollerImpl) Update(pctx echo.Context) error {

	todoID, err := _utils.StrToUint(pctx.Param("todoID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid todoID", nil)
	}

	todoReq := new(_todoModel.TodoReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(todoReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	ctx := pctx.Request().Context()
	_, err = c.todoService.Update(ctx, todoID, todoReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *todoContollerImpl) Delete(pctx echo.Context) error {

	todoID, err := _utils.StrToUint(pctx.Param("todoID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid todoID", nil)
	}

	ctx := pctx.Request().Context()
	_, err = c.todoService.Delete(ctx, todoID)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *todoContollerImpl) FindAll(pctx echo.Context) error {

	TodoSearchReq := new(_todoModel.TodoSearchReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(TodoSearchReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	ctx := pctx.Request().Context()
	todoSearch, err := c.todoService.FindAll(ctx, TodoSearchReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, todoSearch, "", nil)

}
