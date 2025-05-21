package service

import (
	"context"

	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"
)

type ItemService interface {
	Create(ctx context.Context, itemReq *_itemModel.ItemReq) (*_itemModel.ItemRes, error)
	Get(ctx context.Context, itemID uint) (*_itemModel.ItemWithLangRes, error)
	Update(ctx context.Context, itemID uint, itemReq *_itemModel.ItemReq) (*_itemModel.ItemRes, error)
	Delete(ctx context.Context, itemID uint) (*_itemModel.ItemWithLangRes, error)
	FindAll(ctx context.Context, todoSearchReq *_itemModel.ItemSearchReq) (*_itemModel.ItemSearchRes, error)
}
