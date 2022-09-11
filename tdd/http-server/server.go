package main

import (
	"fmt"
	"net/http"
	"strings"
)

func ListenAndServe(addr string, handler Handler) error {
	return nil
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}
