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
	router.Handle("/status", StatusHandler).Methods("GET")
	router.Handle("/cryptos", JwtMiddleware.Handler(CryptoHandler)).Methods("GET")
	router.Handle("/cryptos/{slug}/info", JwtMiddleware.Handler(GetCryptoHandler)).Methods("POST")
	router.Handle("/get-token", GetTokenHandler).Methods("GET")
	err := http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		panic(err)
	}
}
