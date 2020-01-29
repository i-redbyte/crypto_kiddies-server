package app

import (
	u "github.com/ilya-sokolov/crypto_kiddies-server/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Запрашиваемые данные не были найдены на сервере"))
		next.ServeHTTP(w, r)
	})
}
