package exception

import "errors"

type TodoException struct{}

func CannotCreateTodo() error {
	return errors.New("Cannot create todo")
}

func CannotUpdateTodo() error {
	return errors.New("Cannot update todo")
}

func CannotDeleteTodo() error {
	return errors.New("Cannot delete todo")
}

func NotFoundTodo() error {
	return errors.New("Not found todo")
}

func TodoStatusInvalid() error {
	return errors.New("Todostatus invalid")
}