package controllers

import (
	"encoding/json"
	"github.com/ilya-sokolov/crypto_kiddies-server/model"
	u "github.com/ilya-sokolov/crypto_kiddies-server/utils"
	"net/http"
)

var CreateGame = func(w http.ResponseWriter, r *http.Request) {
	game := &model.Game{}
	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Неверный запрос"))
		return
	}
	response := game.CreateGame()
	if response["status"] == false {
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, response)
}
