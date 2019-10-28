package main

import (
	"cryptokiddies-server/app"
	c "cryptokiddies-server/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtMiddleware)
	router.Handle("/", http.FileServer(http.Dir("./html/")))
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.Handle("/api/status", StatusHandler).Methods("GET")
	router.HandleFunc("/api/cryptos", c.GetCryptoAlgorithmsHandler).Methods("GET")
	router.HandleFunc("/api/cryptos/{slug}/info", c.GetCryptoHandler).Methods("GET")
	// TODO: Red_byte release it's dummies:
	router.HandleFunc("/api/cryptos/{slug}/list", c.GetCryptoListHandler).Methods("GET")
	router.HandleFunc("/api/cryptos/{slug}/text", c.GetCryptoTextHandler).Methods("POST")
	router.HandleFunc("/api/cryptos/{slug}/create", c.CreateCryptoTextHandler).Methods("POST")
	router.HandleFunc("/api/cryptos/{slug}/send", c.SendAnswerToChef).Methods("POST")

	router.HandleFunc("/api/user/registration", c.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", c.Authenticate).Methods("POST")
	router.HandleFunc("/api/game/new", c.CreateGame).Methods("POST")
	//************ test html
	router.HandleFunc("/login", GetLogin).Methods("GET")
	router.HandleFunc("/login", PostLogin).Methods("POST")
	err := http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		panic(err)
	}
}
