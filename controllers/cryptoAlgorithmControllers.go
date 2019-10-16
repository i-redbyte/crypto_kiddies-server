package controllers

import (
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var GetCryptoAlgorithmsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "success")
	data := model.GetCryptos()
	w.Header().Set("Content-Type", "application/json")
	if len(data) == 0 && data == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Методы шифрования не найдены"))
		return
	}
	response["data"] = data
	u.Respond(w, response)
})

var GetCryptoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var crypto model.Crypto
	vars := mux.Vars(r)
	path := vars["path"]
	crypto = *model.GetCryptoByPath(path)
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
