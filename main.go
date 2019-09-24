package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)
	err := manners.ListenAndServe(":4000", handler)
	if err != nil {
		panic(err)
	}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Crypto Punk"
	}
	_, _ = fmt.Fprintln(res, "Crypto kiddies home page")
	_, _ = fmt.Fprint(res, "Hello, ", name)

}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}
