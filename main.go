package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ilya-sokolov/crypto_kiddies-server/app"
	"github.com/ilya-sokolov/crypto_kiddies-server/appHtml"
	"github.com/ilya-sokolov/crypto_kiddies-server/controllers"

	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtMiddleware)
	router.Handle("/", http.FileServer(http.Dir("./html/")))
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.Handle("/api/status", appHtml.StatusHandler).Methods("GET")
	//-->CRYPTOS
	router.HandleFunc("/api/cryptos", controllers.GetCryptoAlgorithmsHandler).Methods("GET")
	router.HandleFunc("/api/cryptos/{slug}/info", controllers.GetCryptoHandler).Methods("GET")
	router.HandleFunc("/api/cryptos/{slug}/list", controllers.GetCryptoListHandler).Methods("GET")
	router.HandleFunc("/api/cryptos/{slug}/text", controllers.GetCryptoTextHandler).Methods("GET")
	router.HandleFunc("/api/cryptos/{slug}/create", controllers.CreateCryptoTextHandler).Methods("POST")
	router.HandleFunc("/api/cryptos/{slug}/send", controllers.SendAnswerToChef).Methods("POST")
	//-->USER
	router.HandleFunc("/api/user/registration", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/game/new", controllers.CreateGame).Methods("POST")
	//************ test html
	router.HandleFunc("/login", appHtml.GetLogin).Methods("GET")
	router.HandleFunc("/login", appHtml.PostLogin).Methods("POST")
	err := http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		panic(err)
	}
}
