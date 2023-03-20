package Business

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Models"
	"context"
)

type GetPagingItemStorage interface {
	GetPagingItem(
		ctx context.Context,
		filter *Models.Filter,
		paging *Common.Paging,
		moreKeys ...string,
	) ([]Models.TodoItem, error)
}

type GetPagingItemBusiness struct {
	store GetPagingItemStorage
}

func NewGetPagingItemBusiness(store GetPagingItemStorage) *GetPagingItemBusiness {
	return &GetPagingItemBusiness{store: store}
}
func (biz *GetPagingItemBusiness) GetPagingItem(ctx context.Context, filter *Models.Filter, paging *Common.Paging) ([]Models.TodoItem, error) {
	data, err := biz.store.GetPagingItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}
