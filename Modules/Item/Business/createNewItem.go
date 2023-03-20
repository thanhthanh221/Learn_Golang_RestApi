package Business

import (
	"GoLang/Modules/Item/Models"
	"context"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *Models.TodoItemCreation) error
}

type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store: store}
}
func (biz *createItemBusiness) CreateNewItem(ctx context.Context, data *Models.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return Models.ErrTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
