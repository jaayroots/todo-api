package service

import (
	"context"

	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"
)

type ItemService interface {
	Create(ctx context.Context, itemReq *_itemModel.ItemReq) (*_itemModel.ItemRes, error)
	Get(ctx context.Context, todoID uint) (*_itemModel.ItemResV2, error)

}
