package controllers

import (
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Неверный запрос"))
		return
	}
	response := account.Create()
	if response["status"] == false {
		w.Header().Add("Content-Type", "application/json")
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
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, response)
}