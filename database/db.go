package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect(host, database, user, password string) (*gorm.DB, error) {
	var err error
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, password)
	DB, err = gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	return DB, err
}

// Disconnect - close db
func Disconnect(database *gorm.DB) error {
	return database.Close()
}
