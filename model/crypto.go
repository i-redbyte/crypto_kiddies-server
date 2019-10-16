package model

type Crypto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

func GetCryptos() []Crypto {
	var cryptos []Crypto
	GetDB().Table("cryptos").Find(&cryptos)
	return cryptos
}

func GetCryptoByPath(path string) *Crypto {
	crypto := &Crypto{}
	GetDB().Table("cryptos").Where("path = ?", path).First(crypto)
	if crypto.Name == "" {
		return nil
	}
	return crypto
}
