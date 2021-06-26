package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	canceled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true

}

func TestServer(t *testing.T) {
	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, World!"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assert.Equal(t, true, store.canceled)
	})

	t.Run("should not cacnel", func(t *testing.T) {
		data := "Hello, World!"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, _ := context.WithCancel(request.Context())
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		assert.Equal(t, false, store.canceled)
		assert.Equal(t, data, response.Body.String())
	})
}
