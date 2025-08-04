package biz

import (
	"context"
	"to-list/common"
	"to-list/module/item/model"
)

// Handler -> Biz -> Repository -> Storage

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{
		store: store,
	}
}

func (biz *createItemBiz) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
