package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("./html/")))
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.Handle("api/status", StatusHandler).Methods("GET")
	router.Handle("api/cryptos", JwtMiddleware.Handler(CryptoHandler)).Methods("GET")
	router.Handle("api/cryptos/{slug}/info", JwtMiddleware.Handler(GetCryptoHandler)).Methods("POST")
	router.Handle("api/get-token", GetTokenHandler).Methods("GET")
	router.Handle("/login", GetLogin).Methods("GET")
	router.Handle("/login", PostLogin).Methods("POST")
	err := http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		panic(err)
	}
}
