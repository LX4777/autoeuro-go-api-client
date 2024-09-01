package service

import (
	"autoeuro-go-api-client/client"
	"autoeuro-go-api-client/requests_data"
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

// GetWarehouses Список складов, доступных клиенту для способа доставки
func (s *AutoeuroService) GetWarehouses(data requests_data.GetWarehousesRequestData) (*responses.GetWarehousesResponse, error) {
	return client.Request[responses.GetWarehousesResponse](s.ApiClient, "/get_warehouses", data)
}

// GetPayers Получение списка плательщиков (подразделений) для клиента
func (s *AutoeuroService) GetPayers() (*responses.GetPayersResponse, error) {
	return client.Request[responses.GetPayersResponse](s.ApiClient, "/get_payers", nil)
}

// GetBrands Получение списка производителей
func (s *AutoeuroService) GetBrands() (*responses.GetBrandsResponse, error) {
	return client.Request[responses.GetBrandsResponse](s.ApiClient, "/get_brands", nil)
}

// SearchBrands Получение списка брендов, у которых есть искомый артикул
func (s *AutoeuroService) SearchBrands(data requests_data.SearchBrandsRequestData) (*responses.SearchBrandsResponse, error) {
	return client.Request[responses.SearchBrandsResponse](s.ApiClient, "/search_brands", data)
}

// SearchItems Поиск доступных для заказа товаров по бренду и коду с кроссами из наличия и под заказ
func (s *AutoeuroService) SearchItems(data requests_data.SearchItemsRequestData) (*responses.SearchItemsResponse, error) {
	return client.Request[responses.SearchItemsResponse](s.ApiClient, "/search_items", data)
}