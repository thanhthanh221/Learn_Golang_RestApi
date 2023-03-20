package Storange

import "gorm.io/gorm"

type mySqlStore struct {
	db *gorm.DB
}

func NewSQLSTore(db *gorm.DB) *mySqlStore {
	return &mySqlStore{db: db}
}
