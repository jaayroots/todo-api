package service

import (
	"context"

	_itemMapper "github.com/jaayroots/todo-api/pkg/item/mapper"
	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"
	_itemRepository "github.com/jaayroots/todo-api/pkg/item/repository"
	_userRepository "github.com/jaayroots/todo-api/pkg/user/repository"

	_utils "github.com/jaayroots/todo-api/utils"
)

type itemServiceImpl struct {
	itemRepository _itemRepository.ItemRepository
	userRepository _userRepository.UserRepository
}

func NewItemServiceImpl(
	itemRepository _itemRepository.ItemRepository,
	userRepository _userRepository.UserRepository,
) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
		userRepository: userRepository,
	}
}

func (s *itemServiceImpl) Create(ctx context.Context, itemReq *_itemModel.ItemReq) (*_itemModel.ItemRes, error) {

	itemEntity, err := _itemMapper.ToItemEntity(itemReq, 0)
	if err != nil {
		return nil, err
	}

	itemEntity, err = s.itemRepository.Create(ctx, itemEntity)
	if err != nil {
		return nil, err
	}

	userIDArray := _utils.ExtractAuditUserID(itemEntity)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	itemRes := _itemMapper.ToItemRes(ctx, itemEntity, userArr)
	return itemRes, nil
}

func (s *itemServiceImpl) Get(ctx context.Context, itemID uint) (*_itemModel.ItemWithLangRes, error) {

	itemEntity, err := s.itemRepository.FindByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	userIDArray := _utils.ExtractAuditUserID(itemEntity)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	itemRes := _itemMapper.ToItemWithLangRes(ctx, itemEntity, userArr)
	return itemRes, nil
}

func (s *itemServiceImpl) Update(ctx context.Context, itemID uint, itemReq *_itemModel.ItemReq) (*_itemModel.ItemRes, error) {

	itemEntity, err := _itemMapper.ToItemEntity(itemReq, itemID)
	if err != nil {
		return nil, err
	}

	_utils.DumpAndExit(itemEntity)

	itemEntity, err = s.itemRepository.Update(ctx, itemEntity)
	if err != nil {
		return nil, err
	}

	userIDArray := _utils.ExtractAuditUserID(itemEntity)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	itemRes := _itemMapper.ToItemRes(ctx, itemEntity, userArr)
	return itemRes, nil
}

func (s *itemServiceImpl) Delete(ctx context.Context, itemID uint) (*_itemModel.ItemWithLangRes, error) {

	itemEntity, err := s.itemRepository.Delete(ctx, itemID)
	if err != nil {
		return nil, err
	}

	userIDArray := _utils.ExtractAuditUserID(itemEntity)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	itemRes := _itemMapper.ToItemWithLangRes(ctx, itemEntity, userArr)
	return itemRes, nil

}
