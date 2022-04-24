package main

import (
	"fmt"
	"net/http"
	"time"
	"testing"
)

type Store interface {
	Fetch() string
	Cancel()
}
type SpyStore struct {
	response string
	cancelled bool
	t *testing.T
}

 func (s *SpyStore) Fetch() string{
	 time.Sleep(100 * time.Millisecond)
	 return s.response
 }

func (s *SpyStore) Cancel(){
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled(){
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")

	}
}

func (s *SpyStore) assertWasNotCancelled(){
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
