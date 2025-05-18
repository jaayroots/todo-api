package repository

import (
	"context"
	"errors"

	databases "github.com/jaayroots/todo-api/database"
	"github.com/jaayroots/todo-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_todoException "github.com/jaayroots/todo-api/pkg/todo/exception"
)

type todoRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewTodoRepositoryImpl(db databases.Database, logger echo.Logger) TodoRepository {
	return &todoRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *todoRepositoryImpl) Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error) {

	todoEntity := new(entities.Todo)

	err := r.db.Connect().
		WithContext(ctx).
		Create(todo).
		Scan(todoEntity).
		Error

	if err != nil {
		return nil, _todoException.CannotCreateTodo()
	}
	return todo, nil

}

func (r *todoRepositoryImpl) Update(ctx context.Context, todoID int, todo *entities.Todo) (*entities.Todo, error) {

	_, err := r.FindByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	todoEntity := new(entities.Todo)

	err = r.db.Connect().
		Model(&entities.Todo{}).
		Where("id = ?", todoID).
		Updates(todo).
		Scan(todoEntity).
		Error
	if err != nil {
		return nil, _todoException.CannotUpdateTodo()
	}

	return todoEntity, nil
}

func (r *todoRepositoryImpl) Delete(ctx context.Context, todoID int) (*entities.Todo, error) {

	todoEntity, err := r.FindByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	err = r.db.Connect().WithContext(ctx).
		Model(&entities.Todo{}).
		Where("id = ?", todoID).
		Delete(&entities.Todo{}).Error
	if err != nil {
		return nil, _todoException.CannotDeleteTodo()
	}

	return todoEntity, nil
}

func (r *todoRepositoryImpl) FindByID(ctx context.Context, todoID int) (*entities.Todo, error) {

	todo := new(entities.Todo)
	err := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Todo{}).
		First(todo, todoID).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, _todoException.NotFoundTodo()
		}
		return nil, err
	}

	return todo, nil
}
