package model

import (
	. "github.com/ilya-sokolov/crypto_kiddies-server/database"
	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Name      string `json:"name"`
	Algorithm Crypto `json:"algorithm"`
}

func (game *Game) CreateGame() map[string]interface{} {
	DB.Create(game)
	// TODO: Red_byte refactoring this:
	//if game.ID <= 0 {
	//	return u.Message(false, "Не удалось создать игру, ошибка подключения.")
	//}
	//response := u.Message(true, "Новая игра")
	//response["game"] = game
	//return response
	return map[string]interface{}{}
}

func (game *Game) RemoveGame() map[string]interface{} {
	DB.Delete(&game)
	// TODO: Red_byte refactoring this:
	//response := u.Message(true, "Игра "+game.Name+" удалена")
	//return response
	return map[string]interface{}{}
}
