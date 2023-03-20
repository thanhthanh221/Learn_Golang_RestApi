package Storange

import (
	"GoLang/Common"
	"GoLang/Modules/Item/Models"
	"context"
)

func (s *mySqlStore) GetPagingItem(ctx context.Context, filter *Models.Filter, paging *Common.Paging, moreKeys ...string) ([]Models.TodoItem, error) {
	var result []Models.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}
	if err := s.db.Table(Models.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
