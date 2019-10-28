package model

import "github.com/jinzhu/gorm"

type GameText struct {
	gorm.Model
	Text   string `json:"text"`
	Key    string `json:"key"`
	UserId uint   `json:"userId"` // last user?
	// TODO: Red_byte what fields are still needed ?
}

// TODO: Red_byte release it
