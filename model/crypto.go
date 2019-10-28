package model

type Crypto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
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
