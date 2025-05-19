package service

import (
	"context"

	_todoModel "github.com/jaayroots/todo-api/pkg/todo/model"
)

type TodoService interface {
	Get(ctx context.Context, todoID uint) (*_todoModel.TodoRes, error)
	Create(ctx context.Context, todoReq *_todoModel.TodoReq) (*_todoModel.TodoRes, error)
	Update(ctx context.Context, todoID uint, todoReq *_todoModel.TodoReq) (*_todoModel.TodoRes, error)
	Delete(ctx context.Context, todoID uint) (*_todoModel.TodoRes, error)
	FindAll(ctx context.Context, todoSearchReq *_todoModel.TodoSearchReq) (*_todoModel.TodoSearchRes, error)
}
