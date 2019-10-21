package model

type Game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Algorithm   Crypto `json:"algorithm"`
	Description string `json:"description"`
}
