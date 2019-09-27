package main

import (
	"cryptokiddies-server/crypt"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/transparent", transparent)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "crypto punk"
	}
	_, _ = fmt.Fprint(res, "hello, ", name)
}
func transparent(res http.ResponseWriter, req *http.Request) {
	fmt.Println("test")
	query := req.URL.Query()
	key := query.Get("key")
	text := query.Get("text")
	// TODO: Red_byte test data (remove it)
	if text == "" {
		text = "Go и Kotlin - близнецы-братья! Кто более матери истории ценен"
	}
	if key == "" {
		key = "си плюс плюс"
	}
	crypt.SetStringKey(key)
	encryptText := crypt.Encrypt([]rune(text))
	fmt.Println(key, text, encryptText)
	_, _ = fmt.Fprintln(res, "Encrypt text:", encryptText)
	_, _ = fmt.Fprintln(res, "Decrypt text:", crypt.Decrypt([]rune(encryptText)))
}
