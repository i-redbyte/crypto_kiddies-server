package controllers

import (
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"encoding/json"
	"net/http"
)

var CreateGame = func(w http.ResponseWriter, r *http.Request) {
	game := &model.Game{}
	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		u.Respond(w, u.Message(false, "Неверный запрос"))
		return
	}
	response := game.CreateGame()
	if response["status"] == false {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, response)
}