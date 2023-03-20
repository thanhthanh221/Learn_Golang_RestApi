package Business

import (
	"GoLang/Modules/Item/Models"
	"context"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*Models.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *Models.TodoItemUpdate) error
}

type UpdateItemBusiness struct {
	store UpdateItemStorage
}

func NewGetUpdateItemBusiness(store UpdateItemStorage) *UpdateItemBusiness {
	return &UpdateItemBusiness{store: store}
}
func (biz *UpdateItemBusiness) UpdateItemById(ctx context.Context, id int, dataUpdate *Models.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == Models.ItemStatusDelete {
		return Models.ErrItemDeleted
	}
	if biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
