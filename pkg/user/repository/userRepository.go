package repository

import "github.com/jaayroots/todo-api/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	GetByUserID(userID int) (*entities.User, error)
	Update(userID int, user *entities.User) (*entities.User, error)
	Delete(userID int) error
}
