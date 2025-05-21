package repository

import (
	"context"
	"errors"

	databases "github.com/jaayroots/todo-api/database"
	"github.com/jaayroots/todo-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemException "github.com/jaayroots/todo-api/pkg/item/exception"

	_utils "github.com/jaayroots/todo-api/utils"

	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"
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
		return nil, _itemException.CannotCreateItem()
	}
	return item, nil

}

func (r *itemRepositoryImpl) FindByID(ctx context.Context, itemID uint) (*entities.Item, error) {

	item := new(entities.Item)
	err := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Item{}).
		Preload("Translations").
		First(item, itemID).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, _itemException.NotFoundItem()
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

func (r *itemRepositoryImpl) FindOneTranslationByID(ctx context.Context, itemID uint) (*entities.ItemTranslation, error) {

	itemTranslation := new(entities.ItemTranslation)
	err := r.db.Connect().
		Where("id = ?", itemID).
		Find(&itemTranslation).
		First(itemTranslation, itemID).
		Error

	if err != nil {
		return nil, err
	}

	return itemTranslation, nil
}

func (r *itemRepositoryImpl) Update(ctx context.Context, itemEntity *entities.Item) (*entities.Item, error) {

	item, err := r.FindByID(ctx, itemEntity.ID)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, _itemException.CannotUpdateItem()
	}

	itemTranslationEntity, err := r.FindTranslationByID(ctx, item.ID)
	if err != nil {
		return nil, err
	}

	for _, itemTranslation := range itemTranslationEntity {
		r.DeleteTranslation(ctx, itemTranslation.ID)
	}

	scanResult := new(entities.Item)
	err = r.db.Connect().WithContext(ctx).
		Updates(itemEntity).
		Scan(scanResult).
		Error
	if err != nil {
		return nil, nil
	}

	return scanResult, nil

}

func (r *itemRepositoryImpl) DeleteTranslation(ctx context.Context, itemID uint) (*entities.ItemTranslation, error) {

	itemTranslationEntity, err := r.FindOneTranslationByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	err = r.db.Connect().
		WithContext(ctx).
		Delete(itemTranslationEntity).Error

	if err != nil {
		return nil, nil
	}

	return itemTranslationEntity, nil
}

func (r *itemRepositoryImpl) Delete(ctx context.Context, itemID uint) (*entities.Item, error) {

	item, err := r.FindByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	itemTranslationEntity, err := r.FindTranslationByID(ctx, item.ID)
	if err != nil {
		return nil, err
	}

	for _, itemTranslation := range itemTranslationEntity {
		r.DeleteTranslation(ctx, itemTranslation.ID)
	}

	err = r.db.Connect().
		WithContext(ctx).
		Delete(item).Error

	if err != nil {
		return nil, nil
	}

	return item, nil
}

func (r *itemRepositoryImpl) FindAll(ctx context.Context, itemSearchReq *_itemModel.ItemSearchReq) ([]*entities.Item, int, error) {

	var item []*entities.Item
	var total int64

	query := r.db.Connect().WithContext(ctx).
		Model(&entities.Item{})

	offset, limit, _ := _utils.PaginateCalculate(itemSearchReq.Page, itemSearchReq.Limit, 0)
	query = r.searchFilter(query, itemSearchReq)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).
		Preload("Translations").
		Order("created_at DESC").
		Find(&item).Error; err != nil {
		return nil, 0, err
	}

	return item, int(total), nil
}

func (r *itemRepositoryImpl) searchFilter(query *gorm.DB, itemSearchReq *_itemModel.ItemSearchReq) *gorm.DB {

	query = r.filterTitle(query, itemSearchReq.Filter)
	query = r.filterDescription(query, itemSearchReq.Filter)
	return query
}

func (r *itemRepositoryImpl) filterTitle(query *gorm.DB, todoFilterReq _itemModel.ItemFilterReq) *gorm.DB {

	title := todoFilterReq.Title
	if title == nil {
		return query
	}

	subQuery := query.Session(&gorm.Session{}).Model(&entities.ItemTranslation{}).
		Select("item_id").
		Where("title LIKE ?", "%"+*title+"%")

	query = query.Where("id IN (?)", subQuery)

	return query
}

func (r *itemRepositoryImpl) filterDescription(query *gorm.DB, todoFilterReq _itemModel.ItemFilterReq) *gorm.DB {

	description := todoFilterReq.Description
	if description == nil {
		return query
	}

	subQuery := query.Session(&gorm.Session{}).Model(&entities.ItemTranslation{}).
		Select("item_id").
		Where("description LIKE ?", "%"+*description+"%")

	query = query.Where("id IN (?)", subQuery)

	return query
}
