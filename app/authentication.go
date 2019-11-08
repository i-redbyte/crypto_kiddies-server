package app

import (
	"context"
	"cryptokiddies-server/model"
	u "cryptokiddies-server/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JwtMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{
			"/api/user/registration",
			"/api/user/login",
			"/login",
			"/api/status",
		}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" { //return code 403  Unauthorized
			response = u.Message(false, "Отсутствует токен авторизации")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //check `Bearer {token-body}`
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //get jwt token
		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //return code 403  Unauthorized
			response = u.Message(false, "Неверно сформированный токен аутентификации")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Невалидный токен")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}
		_ = fmt.Sprintf("User %d", tk.UserId) //for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
