package service

import (
	"autoeuro-go-api-client/client"
	"autoeuro-go-api-client/responses"
)

type AutoeuroService struct {
	*client.ApiClient
}

func NewAutoeuroService(config client.ApiClientConfig) *AutoeuroService {
	return &AutoeuroService{client.NewApiClient(config)}
}

func (s *AutoeuroService) GetBalance() (*responses.GetBalanceResponse, error) {
	return client.Request[responses.GetBalanceResponse](s.ApiClient, "/get_balance", nil)
}
