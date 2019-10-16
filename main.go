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
	router.Handle("/api/cryptos", c.GetCryptoAlgorithmsHandler).Methods("GET")
	router.Handle("/api/cryptos/{path}/info", c.GetCryptoHandler).Methods("GET")

	router.HandleFunc("/api/user/registration", c.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", c.Authenticate).Methods("POST")
	//************ test html
	router.Handle("/login", GetLogin).Methods("GET")
	router.Handle("/login", PostLogin).Methods("POST")
	err := http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		panic(err)
	}
}
