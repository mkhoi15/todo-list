package storage

import (
	"context"
	"to-list/common"
	"to-list/module/item/model"
)

func (s *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
