package main

import (
	"github.com/ilya-sokolov/crypto_kiddies-server/app/rest/api"
	"github.com/ilya-sokolov/crypto_kiddies-server/database"
	"github.com/ilya-sokolov/crypto_kiddies-server/storage"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	database.DB, err = database.Connect(dbHost, dbName, userName, password)
	storage.InitMigration()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = database.Disconnect(database.DB)
		if err != nil {
			panic(err)
		}
	}()
	app := api.Router()
	err = app.Run(":4000")
	if err != nil {
		panic(err)
	}
}
