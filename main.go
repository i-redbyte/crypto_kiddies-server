package main

import (
	"fmt"
	"net/http"
)

func homePage(response http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(response, req)
		return
	}
	_, err := fmt.Fprint(response, "Crypto kiddies home page")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}
