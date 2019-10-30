package controllers

import (
	"cryptokiddies-server/crypt"
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var GetCryptoAlgorithmsHandler = func(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "success")
	data := model.GetCryptos()
	if len(data) == 0 && data == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Методы шифрования не найдены"))
		return
	}
	response["data"] = data
	u.Respond(w, response)
}

var GetCryptoHandler = func(w http.ResponseWriter, r *http.Request) {
	var crypto model.Crypto
	vars := mux.Vars(r)
	slug := vars["slug"]
	crypto = *model.GetCryptoByPath(slug)
	response := u.Message(true, "success")
	if crypto.Slug != "" {
		response["data"] = crypto
	} else {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Метод шифрования не найден"))
		return
	}
	u.Respond(w, response)
}

var GetCryptoListHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: Red_byte get crypto texts by slug from db
}
var GetCryptoTextHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: Red_byte get crypto text by slug from db
}

var CreateCryptoTextHandler = func(w http.ResponseWriter, r *http.Request) {
	var crypto model.Crypto
	gameText := &model.GameText{}
	err := json.NewDecoder(r.Body).Decode(gameText)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Неверный запрос"))
		return
	}
	vars := mux.Vars(r)
	slug := vars["slug"]
	text := gameText.Text
	key := gameText.Key

	crypto = *model.GetCryptoByPath(slug)
	encryptText, err := crypt.GetCryptoText(slug, text, key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	gameText.AlgorithmName = crypto.Name
	gameText.AlgorithmId = crypto.Id
	gameText.Text = *encryptText
	response := gameText.CreateGameText()

	if response["status"] == false {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, response)
}

var SendAnswerToChef = func(w http.ResponseWriter, r *http.Request) {
	// TODO: Red_byte send an answer to the chef
}
