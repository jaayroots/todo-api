package mapper

import (
	"time"

	"github.com/jaayroots/todo-api/entities"
	"github.com/jaayroots/todo-api/enums"
	_todoException "github.com/jaayroots/todo-api/pkg/todo/exception"
	_todoModel "github.com/jaayroots/todo-api/pkg/todo/model"
)

func ToTodoEntity(todoReq *_todoModel.TodoReq) (*entities.Todo, error) {

	var dueDate *time.Time
	if todoReq.DueDate != 0 {
		dt := time.Unix(todoReq.DueDate, 0)
		dueDate = &dt
	}

	status := enums.TodoStatus(todoReq.Status)
	if !enums.IsValidTodoStatus(int(status)) {
		return nil, _todoException.TodoStatusInvalid()
	}

	todoEntity := &entities.Todo{
		Title:       todoReq.Title,
		Description: todoReq.Description,
		Status:      status,
		DueDate:     dueDate,
	}

	return todoEntity, nil
}

func ToTodoRes(todo *entities.Todo) *_todoModel.TodoRes {

	return &_todoModel.TodoRes{
		ID:          int(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      int(todo.Status),
		DueDate:     int64(todo.DueDate.Unix()),
		CreatedAt:   todo.CreatedAt.Unix(),
		UpdatedAt:   todo.UpdatedAt.Unix(),
	}
}
