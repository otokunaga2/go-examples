package main

import (
	"fmt"
	"net/http"
)

func ListenAndServe(addr string, handler Handler) error {
	return nil
}
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
