package storage

import (
	"context"
	"to-list/common"
	"to-list/module/item/model"
)

func (s *sqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deleteStatus := "Deleted"

	if err := s.db.Model(&model.TodoItem{}).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deleteStatus,
		}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
