package Models

import (
	"GoLang/Common"
	"errors"
)

var (
	ErrTitleIsBlank = errors.New("Title cannot be blank!")
	ErrItemDeleted  = errors.New("Item is deleted!")
)

type TodoItem struct {
	Common.SQLModel
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      *ItemStatus `json:"status"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"colum:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
