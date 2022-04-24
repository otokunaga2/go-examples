package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"context"
	"fmt"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T){
		store := &SpyStore{response: data, t:t}	
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data{
			t.Errorf(`got %s, want %s`, response.Body.String(), data)

		}
		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T){
		store := &SpyStore{response: data, t:t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(15*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		//fmt.Println(store)
		//store.cancelled = true
		svr.ServeHTTP(response, request)

		store.assertWasCancelled()

	})
}

