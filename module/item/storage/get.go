package storage

import (
	"context"
	"gorm.io/gorm"
	"to-list/common"
	"to-list/module/item/model"
)

func (s *sqlStorage) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var item model.TodoItem

	if err := s.db.Where(cond).First(&item).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &item, nil
}
