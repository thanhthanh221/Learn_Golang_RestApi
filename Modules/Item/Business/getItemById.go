package Business

import (
	"GoLang/Modules/Item/Models"
	"context"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*Models.TodoItem, error)
}

type GetItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *GetItemBusiness {
	return &GetItemBusiness{store: store}
}
func (biz *GetItemBusiness) GetItemById(ctx context.Context, id int) (*Models.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
