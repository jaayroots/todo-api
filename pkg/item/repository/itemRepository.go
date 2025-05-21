package repository

import (
	"context"

	"github.com/jaayroots/todo-api/entities"
	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"
)

type ItemRepository interface {
	Create(ctx context.Context, item *entities.Item) (*entities.Item, error)
	FindByID(ctx context.Context, itemID uint) (*entities.Item, error)
	FindTranslationByID(ctx context.Context, itemID uint) ([]*entities.ItemTranslation, error)
	Update(ctx context.Context, item *entities.Item) (*entities.Item, error)
	DeleteTranslation(ctx context.Context, itemID uint) (*entities.ItemTranslation, error)
	Delete(ctx context.Context, itemID uint) (*entities.Item, error)
	FindAll(ctx context.Context, itemSearchReq *_itemModel.ItemSearchReq) ([]*entities.Item, int, error)
}
