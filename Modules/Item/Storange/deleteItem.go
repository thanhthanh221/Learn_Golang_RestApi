package Storange

import (
	"GoLang/Modules/Item/Models"
	"context"
)

func (s *mySqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deleteStatus := Models.ItemStatusDelete
	if err := s.db.Table(Models.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deleteStatus.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}
