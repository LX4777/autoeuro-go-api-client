package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	Server   *httptest.Server
	Response interface{}
}

func NewTestServer(expectedPath string, mockResponse interface{}) *TestServer {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			http.Error(w, "Неожиданный url", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))

	return &TestServer{
		Server:   server,
		Response: mockResponse,
	}
}

func (ts *TestServer) Close() {
	ts.Server.Close()
}
