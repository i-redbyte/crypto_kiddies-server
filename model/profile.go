package model

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
