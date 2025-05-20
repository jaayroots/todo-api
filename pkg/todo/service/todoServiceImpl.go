package service

import (
	"context"

	_todoMapper "github.com/jaayroots/todo-api/pkg/todo/mapper"
	_todoModel "github.com/jaayroots/todo-api/pkg/todo/model"
	_todoRepository "github.com/jaayroots/todo-api/pkg/todo/repository"
)

type todoServiceImpl struct {
	todoRepository _todoRepository.TodoRepository
}

func NewTodoServiceImpl(
	todoRepository _todoRepository.TodoRepository,
) TodoService {
	return &todoServiceImpl{todoRepository}

}
func (s *todoServiceImpl) Get(ctx context.Context, todoID uint) (*_todoModel.TodoRes, error) {
	todo, err := s.todoRepository.FindByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	todoRes := _todoMapper.ToTodoRes(todo)
	return todoRes, nil
}

func (s *todoServiceImpl) Create(ctx context.Context, todoReq *_todoModel.TodoReq) (*_todoModel.TodoRes, error) {

	todoEntity, err := _todoMapper.ToTodoEntity(todoReq, 0)
	if err != nil {
		return nil, err
	}

	_, err = s.todoRepository.Create(ctx, todoEntity)
	if err != nil {
		return nil, err
	}

	todoRes := _todoMapper.ToTodoRes(todoEntity)
	return todoRes, nil
}

func (s *todoServiceImpl) Update(ctx context.Context, todoID uint, todoReq *_todoModel.TodoReq) (*_todoModel.TodoRes, error) {

	todoEntity, err := _todoMapper.ToTodoEntity(todoReq, todoID)
	if err != nil {
		return nil, err
	}

	todoEntity, err = s.todoRepository.Update(ctx, todoEntity)
	if err != nil {
		return nil, err
	}

	todoRes := _todoMapper.ToTodoRes(todoEntity)
	return todoRes, nil
}

func (s *todoServiceImpl) Delete(ctx context.Context, todoID uint) (*_todoModel.TodoRes, error) {

	todoEntity, err := s.todoRepository.Delete(ctx, todoID)
	if err != nil {
		return nil, err
	}

	todoRes := _todoMapper.ToTodoRes(todoEntity)
	return todoRes, nil
}

func (s *todoServiceImpl) FindAll(ctx context.Context, todoSearchReq *_todoModel.TodoSearchReq) (*_todoModel.TodoSearchRes, error) {
	todos, total, err := s.todoRepository.FindAll(ctx, todoSearchReq)
	if err != nil {
		return nil, err
	}

	return _todoMapper.ToTodoSearchRes(todoSearchReq, todos, total), nil

}
