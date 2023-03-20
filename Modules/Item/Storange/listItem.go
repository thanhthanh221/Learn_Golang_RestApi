package Storange

import (
	"GoLang/Modules/Item/Models"
	"context"
)

func (s *mySqlStore) GetListItem(ctx context.Context) ([]Models.TodoItem, error) {
	var result []Models.TodoItem

	if err := s.db.Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
