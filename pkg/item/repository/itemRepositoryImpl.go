package repository

import (
	"context"
	"errors"

	databases "github.com/jaayroots/todo-api/database"
	"github.com/jaayroots/todo-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemRepositoryImpl(db databases.Database, logger echo.Logger) ItemRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemRepositoryImpl) Create(ctx context.Context, item *entities.Item) (*entities.Item, error) {

	itemEntity := new(entities.Item)

	err := r.db.Connect().
		WithContext(ctx).
		Create(item).
		Scan(itemEntity).
		Error

	if err != nil {
		return nil, nil
	}
	return item, nil

}

func (r *itemRepositoryImpl) FindByID(ctx context.Context, itemID uint) (*entities.Item, error) {

	item := new(entities.Item)
	err := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Item{}).
		First(item, itemID).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

func (r *itemRepositoryImpl) FindTranslationByID(ctx context.Context, itemID uint) ([]*entities.ItemTranslation, error) {

	var itemTranslation []*entities.ItemTranslation
	err := r.db.Connect().
		Where("item_id = ?", itemID).
		Find(&itemTranslation).Error

	if err != nil {
		return nil, err
	}

	return itemTranslation, nil
}
