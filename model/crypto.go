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
	AlgorithmSlug string `json:"algorithmSlug"`
	AlgorithmId   uint   `json:"algorithmId"`
	Text          string `json:"text"`
	Key           string `json:"key"`
	CreatorId     uint   `json:"creatorId"`
}

const gtTableName = "game_texts"
const cryptoTableName = "cryptos"

func GetCryptos() []Crypto {
	var cryptos []Crypto
	GetDB().Table(cryptoTableName).Find(&cryptos)
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

func GetGameText(id uint) *GameText {
	gameText := &GameText{}
	GetDB().Table(gtTableName).Where("id = ?", id).First(gameText)
	if gameText.Text == "" {
		return nil
	}
	return gameText
}

func GetGameTexts(slug string) []GameText {
	var texts []GameText
	GetDB().Table(gtTableName).Where("algorithm_slug = ?", slug).Find(&texts)
	fmt.Println(texts)
	return texts
}
