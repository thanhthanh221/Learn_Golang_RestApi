package Storange

import (
	"GoLang/Modules/Item/Models"
	"context"
)

func (s *mySqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *Models.TodoItemUpdate) error {

	if err := s.db.Where(cond).Updates(&dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
