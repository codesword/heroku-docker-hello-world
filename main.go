package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: &myHandler{},
	}

	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/hello" {
		hello(w, r)
		return
	}

	io.WriteString(w, "Route with URL: "+r.URL.String())
}
