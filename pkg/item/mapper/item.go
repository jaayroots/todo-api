package mapper

import (
	"context"

	"github.com/jaayroots/todo-api/entities"
	_itemModel "github.com/jaayroots/todo-api/pkg/item/model"

	_utils "github.com/jaayroots/todo-api/utils"

	"github.com/jaayroots/todo-api/pctxkeys"
)

func ToItemEntity(itemReq *_itemModel.ItemReq, itemID uint) (*entities.Item, error) {

	itemTranslation, err := ToItemItemTranslationEntity(itemReq, itemID)
	if err != nil {
		return nil, err
	}

	itemEntity := &entities.Item{
		ID:           itemID,
		Translations: itemTranslation,
	}

	return itemEntity, nil
}

func ToItemItemTranslationEntity(itemReq *_itemModel.ItemReq, itemID uint) ([]*entities.ItemTranslation, error) {
	translations := make([]*entities.ItemTranslation, 0, len(itemReq.Title))

	for lang, title := range itemReq.Title {
		if _utils.IsValidLang(lang) {
			description := ""
			if itemReq.Description != nil {
				if desc, ok := itemReq.Description[lang]; ok {
					description = desc
				}
			}

			translations = append(translations, &entities.ItemTranslation{
				ItemID:      itemID,
				Lang:        lang,
				Title:       title,
				Description: description,
			})
		}
	}

	return translations, nil
}

func ToItemItemTranslationEntityV2(itemReq *_itemModel.ItemReq, itemID uint) ([]*entities.ItemTranslation, error) {
	translations := make([]*entities.ItemTranslation, 0, len(itemReq.Title))

	for lang, title := range itemReq.Title {
		description := ""
		if itemReq.Description != nil {
			if desc, ok := itemReq.Description[lang]; ok {
				description = desc
			}
		}

		translations = append(translations, &entities.ItemTranslation{
			ItemID:      itemID,
			Lang:        lang,
			Title:       title,
			Description: description,
		})
	}

	return translations, nil
}

func ToItemRes(ctx context.Context, item *entities.Item, users []*entities.User) *_itemModel.ItemRes {

	userMap := _utils.MapperByID(users)

	getFullNameByID := func(userID uint) *string {
		if u, ok := userMap[userID]; ok {
			fullName := u.FirstName + " " + u.LastName
			return &fullName
		}
		return nil
	}

	titleMap := make(_itemModel.LocalizedString)
	descMap := make(_itemModel.LocalizedString)

	for _, t := range item.Translations {
		titleMap[t.Lang] = t.Title
		descMap[t.Lang] = t.Description
	}

	return &_itemModel.ItemRes{
		ID:          int(item.ID),
		Title:       titleMap,
		Description: &descMap,
		CreatedAt:   item.CreatedAt.Unix(),
		UpdatedAt:   item.UpdatedAt.Unix(),
		CreatedBy:   getFullNameByID(item.CreatedBy),
		UpdatedBy:   getFullNameByID(item.UpdatedBy),
	}
}

func ToItemWithLangRes(ctx context.Context, item *entities.Item, users []*entities.User) *_itemModel.ItemWithLangRes {

	lang, _ := ctx.Value(pctxkeys.ContextKeyLang).(string)
	userMap := _utils.MapperByID(users)
	getFullNameByID := func(userID uint) *string {
		if u, ok := userMap[userID]; ok {
			fullName := u.FirstName + " " + u.LastName
			return &fullName
		}
		return nil
	}

	var title, desc string
	for _, t := range item.Translations {
		if t.Lang == lang {
			title = t.Title
			desc = t.Description
			break
		}
	}

	return &_itemModel.ItemWithLangRes{
		ID:          int(item.ID),
		Title:       &title,
		Description: &desc,
		CreatedAt:   item.CreatedAt.Unix(),
		UpdatedAt:   item.UpdatedAt.Unix(),
		CreatedBy:   getFullNameByID(item.CreatedBy),
		UpdatedBy:   getFullNameByID(item.UpdatedBy),
	}
}

func ToItemSearchRes(
	ctx context.Context,
	itemSearchReq *_itemModel.ItemSearchReq,
	user []*entities.User,
	items []*entities.Item,
	total int,
) *_itemModel.ItemSearchRes {

	itemResList := make([]*_itemModel.ItemRes, 0, len(items))
	for _, item := range items {
		itemResList = append(itemResList, ToItemRes(ctx, item, user))
	}

	_, _, totalPage := _utils.PaginateCalculate(itemSearchReq.Page, itemSearchReq.Limit, total)
	return &_itemModel.ItemSearchRes{
		Item: itemResList,
		Paginate: _itemModel.PaginateResult{
			Page:      int64(itemSearchReq.Page),
			TotalPage: int64(totalPage),
			Total:     int64(total),
		},
	}
}
