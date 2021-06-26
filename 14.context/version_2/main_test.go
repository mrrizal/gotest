package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "Hello, World!"
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		assert.Equal(t, true, response.written)
	})

	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, World!"
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		assert.Equal(t, false, response.written)
	})
}
