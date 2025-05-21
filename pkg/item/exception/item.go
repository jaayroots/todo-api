package exception

import "errors"

type ItemException struct{}

func IsExistItem() error {
	return errors.New("Item is already")
}

func CannotCreateItem() error {
	return errors.New("Cannot create item")
}

func CannotUpdateItem() error {
	return errors.New("Cannot update item")
}

func CannotDeleteItem() error {
	return errors.New("Cannot delete item")
}

func NotFoundItem() error {
	return errors.New("Not found item")
}
