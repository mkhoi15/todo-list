package storage

import (
	"context"
	"to-list/common"
	"to-list/module/item/model"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
