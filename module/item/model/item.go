package model

import (
	"errors"
	"strings"
	"to-list/common"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
	ErrItemDeleted        = errors.New("item has been deleted")
)

const (
	EntityName = "Item"
)

type TodoItem struct {
	common.SQLModel
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string { return "todo_items" }

func (i *TodoItemCreate) Validate() error {

	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

type TodoItemCreate struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TodoItemCreate) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
