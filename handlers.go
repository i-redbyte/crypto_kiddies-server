package main

// TODO: split the file into logical components
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Crypto struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var cryptos = []Crypto{
	{1, "Перестановочный шифр", "transposition", "Описание шифра перестановки"},
	{2, "Шифр Цезаря", "cipher_caesar", "Описание шифра цезаря"},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("API is running"))
})

var CryptoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("API is running"))
	payload, _ := json.Marshal(cryptos)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(payload)
})

// TODO: Red_byte rename?
var GetCryptoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var crypto Crypto
	vars := mux.Vars(r)
	slug := vars["slug"]
	for _, cry := range cryptos {
		if cry.Slug == slug {
			crypto = cry
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if crypto.Slug != "" {
		payload, _ := json.Marshal(crypto)
		_, _ = w.Write(payload)
	} else {
		_, _ = w.Write([]byte("Метод шифрования не найден"))
	}
})

var GetLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/login.html")
})

var PostLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	// TODO: Red_byte test until session
	fmt.Println(username, password)
	http.ServeFile(w, r, "html/index.html")
})

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Not Implemented"))
	if err != nil {
		panic(err)
	}
})
