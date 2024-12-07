package service

import (
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	Server *httptest.Server
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

	return &TestServer{
		Server: server,
	}
}

func (ts *TestServer) Close() {
	ts.Server.Close()
}
