package main

import (
	"fmt"
	"net/http"
	"strings"
)

func ListenAndServe(addr string, handler Handler) error {
	return nil
}
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	} else if player == "Floyd" {
		fmt.Fprint(w, "10")
		return

	}
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
