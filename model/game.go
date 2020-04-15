package model

import (
	. "github.com/ilya-sokolov/crypto_kiddies-server/database"
	u "github.com/ilya-sokolov/crypto_kiddies-server/utils"
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Name      string `json:"name"`
	Algorithm Crypto `json:"algorithm"`
}

func (game *Game) CreateGame() map[string]interface{} {
	DB.Create(game)
	if game.ID <= 0 {
		return u.Message(false, "Не удалось создать игру, ошибка подключения.")
	}
	response := u.Message(true, "Новая игра")
	response["game"] = game
	return response
}

func (game *Game) RemoveGame() map[string]interface{} {
	DB.Delete(&game)
	response := u.Message(true, "Игра "+game.Name+" удалена")
	return response
}
