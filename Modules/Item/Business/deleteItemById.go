package Business

import (
	"GoLang/Modules/Item/Models"
	"context"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*Models.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type DeleteItemBusiness struct {
	store DeleteItemStorage
}

func NewDeleteItemBusiness(store DeleteItemStorage) *DeleteItemBusiness {
	return &DeleteItemBusiness{store: store}
}
func (biz *DeleteItemBusiness) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != err {
		return err
	}

	if data.Status != nil && *data.Status == Models.ItemStatusDelete {
		return Models.ErrItemDeleted
	}
	if biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
