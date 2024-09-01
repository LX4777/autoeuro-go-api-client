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

// GetBalance Получение детальной информации о состоянии баланса личного счета
func (s *AutoeuroService) GetBalance() (*responses.GetBalanceResponse, error) {
	return client.Request[responses.GetBalanceResponse](s.ApiClient, "/get_balance", nil)
}

// GetDeliveries Получение массива возможных вариантов получения товара для клиента
func (s *AutoeuroService) GetDeliveries() (*responses.GetDeliveriesResponse, error) {
	return client.Request[responses.GetDeliveriesResponse](s.ApiClient, "/get_deliveries", nil)
}
