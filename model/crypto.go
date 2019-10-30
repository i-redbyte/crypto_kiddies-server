package model

import (
	u "cryptokiddies-server/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Crypto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type GameText struct {
	gorm.Model
	AlgorithmName string `json:"algorithmName"`
	AlgorithmId   uint   `json:"algorithmId"`
	Text          string `json:"text"`
	Key           string `json:"key"`
}

func GetCryptos() []Crypto {
	var cryptos []Crypto
	GetDB().Table("cryptos").Find(&cryptos)
	return cryptos
}

func GetCryptoByPath(slug string) *Crypto {
	crypto := &Crypto{}
	GetDB().Table("cryptos").Where("slug = ?", slug).First(crypto)
	if crypto.Name == "" {
		return nil
	}
	return crypto
}

func (gameText *GameText) CreateGameText() map[string]interface{} {
	gameText.AlgorithmId = 0
	if err := GetDB().Create(gameText).Error; err != nil {
		fmt.Println(err)
	}
	if gameText.ID <= 0 {
		return u.Message(false, "Не удалось создать текст, ошибка подключения к БД.")
	}
	response := u.Message(true, "Текст создан")
	response["data"] = gameText
	return response
}
