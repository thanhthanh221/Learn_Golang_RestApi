package Common

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"colum:id;"`
	CreateAt *time.Time `json:"create_at" gorm:"colum:create_at"'`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"colum:create_at"`
}
