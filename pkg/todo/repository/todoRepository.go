package repository

import (
	"context"

	"github.com/jaayroots/todo-api/entities"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error)
	FindByID(ctx context.Context, todoID int) (*entities.Todo, error)
	Update(ctx context.Context, todoID int, todo *entities.Todo) (*entities.Todo, error)
	Delete(ctx context.Context, todoID int) (*entities.Todo, error)
}
