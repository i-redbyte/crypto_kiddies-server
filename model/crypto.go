package model

import "github.com/jinzhu/gorm"

type Crypto struct {
	gorm.Model
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

func GetCryptoByPath(path string) *Crypto {
	crypto := &Crypto{}
	GetDB().Table("cryptos").Where("path = ?", path).First(crypto)
	if crypto.Name == "" {
		return nil
	}
	return crypto
}
