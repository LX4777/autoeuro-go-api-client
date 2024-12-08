package service

import (
	"github.com/LX4777/autoeuro-go-api-client/client"
	"github.com/LX4777/autoeuro-go-api-client/service"
	"net/http"
	"net/http/httptest"
	"time"
)

type TestServer struct {
	Server  *httptest.Server
	Service *service.AutoeuroService
}

func NewTestServer(expectedPath string, json []byte) *TestServer {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			http.Error(w, "Неожиданный url", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(json); err != nil {
			http.Error(w, "ошибка записи json в ответ", http.StatusInternalServerError)
			return
		}
	}))

	s := service.NewAutoeuroService(client.ApiClientConfig{
		BaseURL: server.URL,
		Token:   "test-api-key",
		Timeout: 10 * time.Second,
	})

	return &TestServer{
		Server:  server,
		Service: s,
	}
}

func (ts *TestServer) Close() {
	ts.Server.Close()
}
