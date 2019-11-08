package main

// TODO: split the file into logical components
import (
	"fmt"
	"net/http"
)

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("API is running"))
})

var GetLogin = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/login.html")
}

var PostLogin = func(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	// TODO: Red_byte test until session
	fmt.Println(username, password)
	http.ServeFile(w, r, "html/index.html")
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Not Implemented"))
	if err != nil {
		panic(err)
	}
})
