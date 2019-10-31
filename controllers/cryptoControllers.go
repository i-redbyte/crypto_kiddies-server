package controllers

import (
	"cryptokiddies-server/crypt"
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Id struct {
	ID int `json:"id"`
}

var GetCryptoAlgorithmsHandler = func(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "success")
	data := model.GetCryptos()
	if data == nil || len(data) == 0 {
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
	response := u.Message(true, "success")
	slug := mux.Vars(r)["slug"]
	data := model.GetGameTexts(slug)
	if data == nil || len(data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Шифрованных текстов нет"))
		return
	}
	response["data"] = data
	u.Respond(w, response)
}
var GetCryptoTextHandler = func(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "success")
	id := &Id{}
	err := json.NewDecoder(r.Body).Decode(id)
	fmt.Println("idString", id)
	if err != nil || id.ID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Неверный запрос"))
		return
	}
	gameText := model.GetGameText(uint(id.ID))
	if gameText == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Данные с id = "+string(id.ID)+"не найдены"))
		return
	}
	response["data"] = gameText
	u.Respond(w, response)
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
	userId, err := model.GetUserId(r)
	if err != nil || userId == nil {
		w.WriteHeader(http.StatusBadRequest)
		if err == nil {
			u.Respond(w, u.Message(false, "Не удалось получить id пользователя"))
		} else {
			u.Respond(w, u.Message(false, err.Error()))
		}
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
	gameText.AlgorithmSlug = crypto.Slug
	gameText.AlgorithmId = crypto.Id
	gameText.Text = *encryptText
	gameText.CreatorId = *userId
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
