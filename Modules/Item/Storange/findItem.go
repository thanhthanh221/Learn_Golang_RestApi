package Storange

import (
	"GoLang/Modules/Item/Models"
	"context"
)

func (s *mySqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*Models.TodoItem, error) {
	var data Models.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
