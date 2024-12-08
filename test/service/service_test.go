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
	expData := responses.SearchItem{
		OfferKey:        "2",
		Stock:           33,
		Brand:           "Brand2",
		Code:            "Code2",
		Name:            "Item 2",
		Packing:         2,
		Price:           240.0,
		Currency:        "USD",
		Amount:          3,
		Unit:            "кт",
		Return:          2,
		OrderBefore:     "3",
		DeliveryTime:    "1211",
		DeliveryTimeMax: "2322",
		Rejects:         95.0,
		Dealer:          0,
		WarehouseName:   "WHN2",
		WarehouseKey:    "WHK2",
	}
	if response != nil {
		assert.Equal(t, expData, response.DATA[3])
		cross := uint8(2)
		expData.Price = 240.3
		expData.Cross = &cross
		assert.Equal(t, expData, response.DATA[0])
		expData.Price = 240
		assert.Equal(t, expData, response.DATA[1])
		expData.Cross = nil
		assert.Equal(t, expData, response.DATA[2])
		assert.Equal(t, len(response.DATA), 4, "неожиданное количество элементов")
		assert.IsTypef(t, responses.SearchItemsResponse{}, *response, "тип ответа не соответствует ожидаемому")
	} else {
		assert.Fail(t, "ответ не получен")
	}

	assert.NoError(t, err, "Неожиданная ошибка от SearchItems")
	assert.NotNil(t, response, "Ответ не получен")
}

func TestGetWarehouses(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"warehouse_id": "warehouseID",
				"warehouse_key": "warehouseKEY",
				"warehouse_name": "warehouseNAME"
			},
			{
				"warehouse_id": "warehouseID2",
				"warehouse_key": "warehouseKEY2",
				"warehouse_name": "warehouseNAME2"
			}
		]
	}`

	ts := NewTestServer("/get_warehouses", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.GetWarehouses(requests.GetWarehousesRequestData{})

	assert.NoError(t, err, "Неожиданная ошибка от GetDeliveries")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.GetWarehousesResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "warehouseID", response.DATA[0].WarehouseID, "WarehouseID не соответствует ожидаемому")
	assert.Equal(t, "warehouseKEY", response.DATA[0].WarehouseKey, "WarehouseKEY не соответствует ожидаемому")
	assert.Equal(t, "warehouseNAME", response.DATA[0].WarehouseName, "WarehouseNAME не соответствует ожидаемому")
	assert.Equal(t, "warehouseID2", response.DATA[1].WarehouseID, "WarehouseID не соответствует ожидаемому")
	assert.Equal(t, "warehouseKEY2", response.DATA[1].WarehouseKey, "WarehouseKEY не соответствует ожидаемому")
	assert.Equal(t, "warehouseNAME2", response.DATA[1].WarehouseName, "WarehouseNAME не соответствует ожидаемому")
}
