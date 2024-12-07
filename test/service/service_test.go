package service

import (
	"encoding/json"
	"github.com/LX4777/autoeuro-go-api-client/client/requests"
	"github.com/LX4777/autoeuro-go-api-client/client/responses"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	mockResponse := responses.GetBalanceResponse{
		DATA: []responses.Balance{
			{
				Balance:      500.0,
				Credit:       300.0,
				Ordered:      700.3,
				Reserved:     400.2,
				Limit:        1000.0,
				PayTomorrow:  722.2,
				ShippingFrom: 23.1,
				Active:       1,
			},
		},
	}
	mockJson, marshalErr := json.Marshal(mockResponse)
	if marshalErr != nil {
		t.Errorf("ошибка маршалинга")
	}

	ts := NewTestServer("/get_balance", mockJson)
	defer ts.Close()

	response, err := ts.Service.GetBalance()

	assert.NoError(t, err, "Неожиданная ошибка от GetBalance")
	assert.NotNil(t, response, "Ответ равен nil")
	assert.Equal(t, mockResponse.DATA[0].Balance, response.DATA[0].Balance, "Неожиданное значение balance")
	assert.Equal(t, mockResponse.DATA[0].Credit, response.DATA[0].Credit, "Неожиданное значение credit")
	assert.Equal(t, mockResponse.DATA[0].Ordered, response.DATA[0].Ordered, "Неожиданное значение ordered")
	assert.Equal(t, mockResponse.DATA[0].Reserved, response.DATA[0].Reserved, "Неожиданное значение reserved")
	assert.Equal(t, mockResponse.DATA[0].Limit, response.DATA[0].Limit, "Неожиданное значение limit")
	assert.Equal(t, mockResponse.DATA[0].PayTomorrow, response.DATA[0].PayTomorrow, "Неожиданное значение payTomorrow")
	assert.Equal(t, mockResponse.DATA[0].ShippingFrom, response.DATA[0].ShippingFrom, "Неожиданное значение shippingForm")
	assert.Equal(t, mockResponse.DATA[0].Active, response.DATA[0].Active, "Неожиданное значение active")
}

func TestGetBrands(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"brand": "MANN-FILTER"
			},
			{
				"brand": "MANN"
			}
		]
	}`

	ts := NewTestServer("/get_brands", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.GetBrands()

	assert.NoError(t, err, "Неожиданная ошибка от GetBrands")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не равна 2")
	assert.Equal(t, "MANN-FILTER", response.DATA[0].Brand, "имя бренда не соответствует ожидаемому")
	assert.Equal(t, "MANN", response.DATA[1].Brand, "имя бренда не соответствует ожидаемому")
}

func TestSearchItems(t *testing.T) {
	jsonMock := []byte(`
	{
		"DATA": [
			{
				"offer_key": "2", 
				"stock": 33, 
				"cross": 2, 
				"brand": "Brand2", 
				"code": "Code2", 
				"name": "Item 2", 
				"packing": 2, 
				"price": 240.3, 
				"currency": "USD", 
				"amount": 3, 
				"unit": "кт", 
				"return": 2, 
				"order_before": "3", 
				"delivery_time": "1211", 
				"delivery_time_max": "2322", 
				"rejects": 95.0, 
				"dealer": 0, 
				"warehouse_name": "WHN2", 
				"warehouse_key": "WHK2"
			},
			{
				"offer_key": "2", 
				"stock": 33, 
				"cross": "2", 
				"brand": "Brand2", 
				"code": "Code2", 
				"name": "Item 2", 
				"packing": "2", 
				"price": "240.0", 
				"currency": "USD", 
				"amount": 3, 
				"unit": "кт", 
				"return": "2", 
				"order_before": "3", 
				"delivery_time": "1211", 
				"delivery_time_max": "2322", 
				"rejects": 95, 
				"dealer": 0, 
				"warehouse_name": "WHN2", 
				"warehouse_key": "WHK2"
			},
			{
				"offer_key": "2", 
				"stock": 33, 
				"cross": null, 
				"brand": "Brand2", 
				"code": "Code2", 
				"name": "Item 2", 
				"packing": "2", 
				"price": "240", 
				"currency": "USD", 
				"amount": 3, 
				"unit": "кт", 
				"return": "2", 
				"order_before": "3", 
				"delivery_time": "1211", 
				"delivery_time_max": "2322", 
				"rejects": 95.0, 
				"dealer": 0, 
				"warehouse_name": "WHN2", 
				"warehouse_key": "WHK2"
			},
			{
				"offer_key": "2", 
				"stock": 33, 
				"brand": "Brand2", 
				"code": "Code2", 
				"name": "Item 2", 
				"packing": "2", 
				"price": 240, 
				"currency": "USD", 
				"amount": 3, 
				"unit": "кт", 
				"return": 2, 
				"order_before": "3", 
				"delivery_time": "1211", 
				"delivery_time_max": "2322", 
				"rejects": 95.0, 
				"dealer": 0, 
				"warehouse_name": "WHN2", 
				"warehouse_key": "WHK2"
			}
		]
	}`)

	ts := NewTestServer("/search_items", jsonMock)
	defer ts.Close()

	response, err := ts.Service.SearchItems(requests.SearchItemsRequestData{})

	assert.NoError(t, err, "Неожиданная ошибка от SearchItems")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, len(response.DATA), 4, "неожиданное количество элементов")
	assert.IsTypef(t, responses.SearchItemsResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "2", response.DATA[0].OfferKey, "неожиданное значение OfferKey")
	assert.Equal(t, 33, response.DATA[0].Stock, "неожиданное значение Stock")
	assert.Equal(t, "Brand2", response.DATA[0].Brand, "неожиданное значение Brand")
	assert.Equal(t, "Code2", response.DATA[0].Code, "неожиданное значение Code")
	assert.Equal(t, "Item 2", response.DATA[0].Name, "неожиданное значение Name")
	assert.Equal(t, 2, response.DATA[0].Packing, "неожиданное значение Packing")
	assert.Equal(t, float32(240.3), response.DATA[0].Price, "неожиданное значение Price")
	assert.Equal(t, "USD", response.DATA[0].Currency, "неожиданное значение Currency")
	assert.Equal(t, 3, response.DATA[0].Amount, "неожиданное значение Amount")
	assert.Equal(t, "кт", response.DATA[0].Unit, "неожиданное значение Unit")
	assert.Equal(t, uint8(2), response.DATA[0].Return, "неожиданное значение Return")
	assert.Equal(t, "3", response.DATA[0].OrderBefore, "неожиданное значение OrderBefore")
	assert.Equal(t, "1211", response.DATA[0].DeliveryTime, "неожиданное значение DeliveryTime")
	assert.Equal(t, "2322", response.DATA[0].DeliveryTimeMax, "неожиданное значение DeliveryTimeMax")
	assert.Equal(t, float32(95.0), response.DATA[0].Rejects, "неожиданное значение Rejects")
	assert.Equal(t, 0, response.DATA[0].Dealer, "неожиданное значение Dealer")
	assert.Equal(t, "WHN2", response.DATA[0].WarehouseName, "неожиданное значение WarehouseName")
	assert.Equal(t, "WHK2", response.DATA[0].WarehouseKey, "неожиданное значение WarehouseKey")
}
