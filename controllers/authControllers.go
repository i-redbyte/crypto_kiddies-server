package controllers

import (
	"encoding/json"
	"github.com/ilya-sokolov/crypto_kiddies-server/model"
	u "github.com/ilya-sokolov/crypto_kiddies-server/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Неверный запрос"))
		return
	}
	response := account.CreateAccount()
	if response["status"] == false {
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, response)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Неверный запрос "))
		return
	}
	response := model.Login(account.Email, account.Password)
	if response["status"] == false {
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, response)
}
