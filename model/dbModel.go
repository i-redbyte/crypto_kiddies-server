package model

import (
	"time"
)

type BaseModel struct {
	Id        int64      `json:"id" gorm:"column:id; type:serial; unsigned; primary_key; auto_increment;not null;"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
