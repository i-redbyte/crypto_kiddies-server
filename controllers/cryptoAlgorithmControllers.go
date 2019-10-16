package controllers

import (
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// TODO: Red_byte get from database
var cryptos = []model.Crypto{
	{1, "Перестановочный шифр", "transposition", "Описание шифра перестановки"},
	{2, "Шифр Цезаря", "cipher_caesar", "Описание шифра цезаря"},
}

var GetCryptoAlgorithmsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	resp["data"] = cryptos
	u.Respond(w, resp)
})

var GetCryptoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var crypto model.Crypto
	vars := mux.Vars(r)
	path := vars["path"]
	for _, cry := range cryptos {
		if cry.Path == path {
			crypto = cry
		}
	}
	response := u.Message(true, "success")
	w.Header().Set("Content-Type", "application/json")
	if crypto.Path != "" {
		response["data"] = crypto
	} else {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Метод шифрования не найден"))
		return
	}
	u.Respond(w, response)
})
