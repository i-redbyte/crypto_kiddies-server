package main

// TODO: split the file into logical components
import (
	"context"
	"cryptokiddies-server/models"
	u "cryptokiddies-server/utils"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
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

var JwtMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/user/registration", "/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Получение токена

		if tokenHeader == "" { //return code 403  Unauthorized
			response = u.Message(false, "Отсутствует токен авторизации")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //check `Bearer {token-body}`
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //get jwt token
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //return code 403  Unauthorized
			response = u.Message(false, "Неверно сформированный токен аутентификации")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Невалидный токен")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		_ = fmt.Sprintf("User %d", tk.UserId) //Полезно для мониторинга
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
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
