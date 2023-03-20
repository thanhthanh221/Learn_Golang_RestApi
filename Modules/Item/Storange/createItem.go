package Storange

import (
	"GoLang/Modules/Item/Models"
	"context"
)

func (s *mySqlStore) CreateItem(ctx context.Context, data *Models.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
