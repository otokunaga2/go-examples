package main

import "net/http"

func ListenAndServe(addr string, handler Handler) error {
	return nil
}
func PlayerServer(w http.ResponseWriter, r *http.Request) {}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
