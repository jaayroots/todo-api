package repository

import "github.com/jaayroots/todo-api/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	CheckDuplicateEmail(user *entities.User) (*entities.User, error)
	GetByUserID(userID uint64) (*entities.User, error)
	Update(userID uint64, user *entities.User) (*entities.User, error)
	Delete(userID uint64) error
}
