package main

import (
	"github.com/ilya-sokolov/crypto_kiddies-server/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbAddr := os.Getenv("db_addr")
	database.DB, err = database.Connect(dbAddr, dbName, userName, password)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = database.Disconnect(database.DB)
		if err != nil {
			panic(err)
		}
	}()
	app := routes.Router()
	//var chat = chat.NewWSServer(db.DB)
	//app.Any("/ws", chat.Handler)
	err = app.Run(":4000")
	if err != nil {
		panic(err)
	}
	//router := mux.NewRouter()
	//router.Use(app.JwtMiddleware)
	//router.Handle("/", http.FileServer(http.Dir("./html/")))
	//router.PathPrefix("/static/").Handler(
	//	http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//
	//router.Handle("/api/status", appHtml.StatusHandler).Methods("GET")
	////-->CRYPTOS
	//router.HandleFunc("/api/cryptos", controllers.GetCryptoAlgorithmsHandler).Methods("GET")
	//router.HandleFunc("/api/cryptos/{slug}/info", controllers.GetCryptoHandler).Methods("GET")
	//router.HandleFunc("/api/cryptos/{slug}/list", controllers.GetCryptoListHandler).Methods("GET")
	//router.HandleFunc("/api/cryptos/{slug}/text", controllers.GetCryptoTextHandler).Methods("GET")
	//router.HandleFunc("/api/cryptos/{slug}/create", controllers.CreateCryptoTextHandler).Methods("POST")
	//router.HandleFunc("/api/cryptos/{slug}/send", controllers.SendAnswerToChef).Methods("POST")
	////-->USER
	//router.HandleFunc("/api/user/registration", controllers.CreateAccount).Methods("POST")
	//router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	//router.HandleFunc("/api/game/new", controllers.CreateGame).Methods("POST")
	////************ test html
	//router.HandleFunc("/login", appHtml.GetLogin).Methods("GET")
	//router.HandleFunc("/login", appHtml.PostLogin).Methods("POST")
	//err := http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, router))
	//if err != nil {
	//	panic(err)
	//}
}
