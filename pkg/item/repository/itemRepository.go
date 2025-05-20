package repository

import (
	"context"

	"github.com/jaayroots/todo-api/entities"
)

type ItemRepository interface {
	Create(ctx context.Context, item *entities.Item) (*entities.Item, error)
	FindByID(ctx context.Context, itemID uint) (*entities.Item, error)
	FindTranslationByID(ctx context.Context, itemID uint) ([]*entities.ItemTranslation, error)
}
