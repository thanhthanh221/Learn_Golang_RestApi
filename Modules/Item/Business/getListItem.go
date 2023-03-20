package Business

import (
	"GoLang/Modules/Item/Models"
	"context"
)

type GetListItemStorage interface {
	GetListItem(ctx context.Context) ([]Models.TodoItem, error)
}

type GetListItemBusiness struct {
	store GetListItemStorage
}

func NewGetListItemBusiness(store GetListItemStorage) *GetListItemBusiness {
	return &GetListItemBusiness{store: store}
}

func (biz *GetListItemBusiness) GetListItem(ctx context.Context) ([]Models.TodoItem, error) {
	data, err := biz.store.GetListItem(ctx)

	if err != nil {
		return nil, err
	}
	return data, nil
}
