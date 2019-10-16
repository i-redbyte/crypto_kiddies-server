package controllers

import (
	"cryptokiddies-server/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Crypto struct {
	Id          int
	Name        string
	Path        string
	Description string
}

// TODO: Red_byte get from database
var cryptos = []Crypto{
	{1, "Перестановочный шифр", "transposition", "Описание шифра перестановки"},
	{2, "Шифр Цезаря", "cipher_caesar", "Описание шифра цезаря"},
}

var GetCryptoAlgorithmsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "success")
	resp["data"] = cryptos
	utils.Respond(w, resp)
})

var GetCryptoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var crypto Crypto
	vars := mux.Vars(r)
	slug := vars["path"]
	for _, cry := range cryptos {
		if cry.Path == slug {
			crypto = cry
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if crypto.Path != "" {
		payload, _ := json.Marshal(crypto)
		_, _ = w.Write(payload)
	} else {
		_, _ = w.Write([]byte("Метод шифрования не найден"))
	}
})
