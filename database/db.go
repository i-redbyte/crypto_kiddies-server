package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// TODO: 30.03.2020 Create db interface
func Connect(host, db, user, password string) (*gorm.DB, error) {
	var err error
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, db, password)
	DB, err = gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	return DB, err
}

// Disconnect - close db
func Disconnect(db *gorm.DB) error {
	return db.Close()
}
