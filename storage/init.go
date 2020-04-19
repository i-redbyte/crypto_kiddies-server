package storage

import . "github.com/ilya-sokolov/crypto_kiddies-server/database"

func InitMigration() {
	DB.AutoMigrate(
		&Account{},
	)
}
