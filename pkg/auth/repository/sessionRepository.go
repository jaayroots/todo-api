package repository

import "github.com/jaayroots/todo-api/entities"

type SessionRepository interface {
	Create(session *entities.Session) (*entities.Session, error)
	Delete(userId uint64) error
	FindByToken(token string) (*entities.Session, error)
}
