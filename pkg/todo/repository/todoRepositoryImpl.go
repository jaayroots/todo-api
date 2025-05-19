package repository

import (
	"context"
	"errors"

	databases "github.com/jaayroots/todo-api/database"
	"github.com/jaayroots/todo-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_todoException "github.com/jaayroots/todo-api/pkg/todo/exception"
	_todoModel "github.com/jaayroots/todo-api/pkg/todo/model"

	_utils "github.com/jaayroots/todo-api/utils"
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

func (r *todoRepositoryImpl) Update(ctx context.Context, todoID uint, todo *entities.Todo) (*entities.Todo, error) {

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

func (r *todoRepositoryImpl) Delete(ctx context.Context, todoID uint) (*entities.Todo, error) {

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

func (r *todoRepositoryImpl) FindByID(ctx context.Context, todoID uint) (*entities.Todo, error) {

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

func (r *todoRepositoryImpl) FindAll(ctx context.Context, todoSearchReq *_todoModel.TodoSearchReq) ([]*entities.Todo, int, error) {

	var todo []*entities.Todo
	var total int64

	query := r.db.Connect().WithContext(ctx).Model(&entities.Todo{})

	offset, limit, _ := _utils.PaginateCalculate(todoSearchReq.Page, todoSearchReq.Limit, 0)
	query = r.searchFilter(query, todoSearchReq)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&todo).Error; err != nil {
		return nil, 0, err
	}

	return todo, int(total), nil
}

func (r *todoRepositoryImpl) searchFilter(query *gorm.DB, todoSearchReq *_todoModel.TodoSearchReq) *gorm.DB {

	query = r.filterTitle(query, todoSearchReq.Filter)
	query = r.filterDescription(query, todoSearchReq.Filter)
	query = r.filterStatus(query, todoSearchReq.Filter)
	return query
}

func (r *todoRepositoryImpl) filterTitle(query *gorm.DB, todoFilterReq _todoModel.TodoFilterReq) *gorm.DB {

	title := todoFilterReq.Title
	if title == nil {
		return query
	}

	query = query.Where("title LIKE ?", "%"+*title+"%")
	return query
}

func (r *todoRepositoryImpl) filterDescription(query *gorm.DB, todoFilterReq _todoModel.TodoFilterReq) *gorm.DB {

	description := todoFilterReq.Description
	if description == nil {
		return query
	}

	query = query.Where("description LIKE ?", "%"+*description+"%")
	return query
}

func (r *todoRepositoryImpl) filterStatus(query *gorm.DB, todoFilterReq _todoModel.TodoFilterReq) *gorm.DB {

	status := todoFilterReq.Status
	if status == nil {
		return query
	}

	query = query.Where("status = ?", *status)
	return query
}
