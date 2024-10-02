// Пакет реализует API Client для работы с веб-сервисом компании АО "Авто-Евро"
//
// Автор: Rozhkov Vasiliy <profilrv@gmail.com> (GitHub: https://github.com/LX4777)
// Лицензия: MIT (https://opensource.org/licenses/MIT)
//
// Данный пакет реализует работу со всеми функциями веб-сервиса на дату 01.09.2024
//
// Для дополнительной информации посетите репозиторий: https://github.com/LX4777/autoeuro-go-api-client

package service

import (
	c "github.com/LX4777/autoeuro-go-api-client/client"
	req "github.com/LX4777/autoeuro-go-api-client/client/requests"
	res "github.com/LX4777/autoeuro-go-api-client/client/responses"
)

type AutoeuroService struct {
	*c.ApiClient
}

func NewAutoeuroService(config c.ApiClientConfig) *AutoeuroService {
	return &AutoeuroService{c.NewApiClient(config)}
}

// GetBalance Получение детальной информации о состоянии баланса личного счета
func (s *AutoeuroService) GetBalance() (*res.GetBalanceResponse, error) {
	return c.Request[res.GetBalanceResponse](s.ApiClient, "/get_balance", nil)
}

// GetDeliveries Получение массива возможных вариантов получения товара для клиента
func (s *AutoeuroService) GetDeliveries() (*res.GetDeliveriesResponse, error) {
	return c.Request[res.GetDeliveriesResponse](s.ApiClient, "/get_deliveries", nil)
}

// GetWarehouses Список складов, доступных клиенту для способа доставки
func (s *AutoeuroService) GetWarehouses(data req.GetWarehousesRequestData) (*res.GetWarehousesResponse, error) {
	return c.Request[res.GetWarehousesResponse](s.ApiClient, "/get_warehouses", data)
}

// GetPayers Получение списка плательщиков (подразделений) для клиента
func (s *AutoeuroService) GetPayers() (*res.GetPayersResponse, error) {
	return c.Request[res.GetPayersResponse](s.ApiClient, "/get_payers", nil)
}

// GetBrands Получение списка производителей
func (s *AutoeuroService) GetBrands() (*res.GetBrandsResponse, error) {
	return c.Request[res.GetBrandsResponse](s.ApiClient, "/get_brands", nil)
}

// SearchBrands Получение списка брендов, у которых есть искомый артикул
func (s *AutoeuroService) SearchBrands(data req.SearchBrandsRequestData) (*res.SearchBrandsResponse, error) {
	return c.Request[res.SearchBrandsResponse](s.ApiClient, "/search_brands", data)
}

// SearchItems Поиск доступных для заказа товаров по бренду и коду с кроссами из наличия и под заказ
func (s *AutoeuroService) SearchItems(data req.SearchItemsRequestData) (*res.SearchItemsResponse, error) {
	return c.Request[res.SearchItemsResponse](s.ApiClient, "/search_items", data)
}

// CreateOrder Оформление заказа
func (s *AutoeuroService) CreateOrder(data req.CreateOrderRequestData) (*res.CreateOrderResponse, error) {
	return c.Request[res.CreateOrderResponse](s.ApiClient, "/create_order", data)
}

// GetOrders Получение списка текущих, завершенных и отменных товаров с деталями заказа
func (s *AutoeuroService) GetOrders(data req.GetOrdersRequestData) (*res.GetOrdersResponse, error) {
	return c.Request[res.GetOrdersResponse](s.ApiClient, "/get_orders", data)
}

// GetStatuses Список возможных статусов для списка заказов
func (s *AutoeuroService) GetStatuses() (*res.GetStatusesResponse, error) {
	return c.Request[res.GetStatusesResponse](s.ApiClient, "/get_statuses", nil)
}
