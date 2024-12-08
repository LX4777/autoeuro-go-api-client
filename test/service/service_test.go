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

func TestGetDeliveries(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"delivery_key": "del1",
				"name": "name1",
				"time_shift_msk": "4"
			},
			{
				"delivery_key": "del2",
				"name": "name2",
				"time_shift_msk": 2 
			}
		]
	}`

	ts := NewTestServer("/get_deliveries", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.GetDeliveries()

	assert.NoError(t, err, "Неожиданная ошибка")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.GetDeliveriesResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "del1", response.DATA[0].DeliveryKey, "поле не соответствует ожидаемому")
	assert.Equal(t, "name1", response.DATA[0].Name, "поле не соответствует ожидаемому")
	assert.Equal(t, uint8(4), response.DATA[0].TimeShiftMSK, "поле не соответствует ожидаемому")
	assert.Equal(t, "del2", response.DATA[1].DeliveryKey, "поле не соответствует ожидаемому")
	assert.Equal(t, "name2", response.DATA[1].Name, "поле не соответствует ожидаемому")
	assert.Equal(t, uint8(2), response.DATA[1].TimeShiftMSK, "поле не соответствует ожидаемому")
}

func TestGetPayers(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"payer_name": "name1",
				"payer_key": "key1"
			},
			{
				"payer_name": "name2",
				"payer_key": "key2"
			}
		]
	}`

	ts := NewTestServer("/get_payers", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.GetPayers()

	assert.NoError(t, err, "Неожиданная ошибка")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.GetPayersResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "name1", response.DATA[0].PayerName, "поле не соответствует ожидаемому")
	assert.Equal(t, "key1", response.DATA[0].PayerKey, "поле не соответствует ожидаемому")
	assert.Equal(t, "name2", response.DATA[1].PayerName, "поле не соответствует ожидаемому")
	assert.Equal(t, "key2", response.DATA[1].PayerKey, "поле не соответствует ожидаемому")
}

func TestSearchBrands(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"brand": "brand1",
				"code": "code1",
				"name": "name1"
			},
			{
				"brand": "brand2",
				"code": "code2",
				"name": "name2"
			}
		]
	}`

	ts := NewTestServer("/search_brands", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.SearchBrands(requests.SearchBrandsRequestData{})

	assert.NoError(t, err, "Неожиданная ошибка")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.SearchBrandsResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "brand1", response.DATA[0].Brand, "поле не соответствует ожидаемому")
	assert.Equal(t, "code1", response.DATA[0].Code, "поле не соответствует ожидаемому")
	assert.Equal(t, "name1", response.DATA[0].Name, "поле не соответствует ожидаемому")
	assert.Equal(t, "brand2", response.DATA[1].Brand, "поле не соответствует ожидаемому")
	assert.Equal(t, "code2", response.DATA[1].Code, "поле не соответствует ожидаемому")
	assert.Equal(t, "name2", response.DATA[1].Name, "поле не соответствует ожидаемому")
}

func TestCreateOrder(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"order_id": 1,
				"result": true,
				"result_description": "res1"
			},
			{
				"order_id": 2,
				"result": false,
				"result_description": "res2"
			}
		]
	}`

	ts := NewTestServer("/create_order", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.CreateOrder(requests.CreateOrderRequestData{})

	assert.NoError(t, err, "Неожиданная ошибка")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.CreateOrderResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, 1, response.DATA[0].OrderID, "поле не соответствует ожидаемому")
	assert.Equal(t, true, response.DATA[0].Result, "поле не соответствует ожидаемому")
	assert.Equal(t, "res1", response.DATA[0].ResultDescription, "поле не соответствует ожидаемому")
	assert.Equal(t, 2, response.DATA[1].OrderID, "поле не соответствует ожидаемому")
	assert.Equal(t, false, response.DATA[1].Result, "поле не соответствует ожидаемому")
	assert.Equal(t, "res2", response.DATA[1].ResultDescription, "поле не соответствует ожидаемому")
}

func TestGetOrders(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
				"brand": "ExampleBrand",
				"code": "EX12345",
				"name": "Example Product",
				"price": 123.45,
				"amount": 10,
				"unit": "pcs",
				"dealer": 1,
				"cancelable": 0,
				"returnable": 1,
				"status_id": 2,
				"status": "In Progress",
				"document": "INV-2024-001",
				"order_id": 101,
				"comment": "Urgent delivery requested.",
				"united": 0,
				"order_date": "2024-12-08T10:30:00Z",
				"order_number": "ORD123456",
				"delivery": "Standard Shipping",
				"delivery_date": "2024-12-15T16:00:00Z",
				"order_key": "abc123xyz"
			}
		]
	}`

	ts := NewTestServer("/get_orders", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.GetOrders(requests.GetOrdersRequestData{})

	assert.NoError(t, err, "Неожиданная ошибка")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 1, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.GetOrdersResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "ExampleBrand", response.DATA[0].Brand, "поле не соответствует ожидаемому")
	assert.Equal(t, "EX12345", response.DATA[0].Code, "поле не соответствует ожидаемому")
	assert.Equal(t, "Example Product", response.DATA[0].Name, "поле не соответствует ожидаемому")
	assert.Equal(t, 123.45, response.DATA[0].Price, "поле не соответствует ожидаемому")
	assert.Equal(t, 10, response.DATA[0].Amount, "поле не соответствует ожидаемому")
	assert.Equal(t, "pcs", response.DATA[0].Unit, "поле не соответствует ожидаемому")
	assert.Equal(t, 1, response.DATA[0].Dealer, "поле не соответствует ожидаемому")
	assert.Equal(t, 0, response.DATA[0].Cancelable, "поле не соответствует ожидаемому")
	assert.Equal(t, 1, response.DATA[0].Returnable, "поле не соответствует ожидаемому")
	assert.Equal(t, 2, response.DATA[0].StatusID, "поле не соответствует ожидаемому")
	assert.Equal(t, "In Progress", response.DATA[0].Status, "поле не соответствует ожидаемому")
	assert.Equal(t, "INV-2024-001", response.DATA[0].Document, "поле не соответствует ожидаемому")
	assert.Equal(t, 101, response.DATA[0].OrderID, "поле не соответствует ожидаемому")
	assert.Equal(t, "Urgent delivery requested.", response.DATA[0].Comment, "поле не соответствует ожидаемому")
	assert.Equal(t, 0, response.DATA[0].United, "поле не соответствует ожидаемому")
	assert.Equal(t, "2024-12-08T10:30:00Z", response.DATA[0].OrderDate, "поле не соответствует ожидаемому")
	assert.Equal(t, "ORD123456", response.DATA[0].OrderNumber, "поле не соответствует ожидаемому")
	assert.Equal(t, "Standard Shipping", response.DATA[0].Delivery, "поле не соответствует ожидаемому")
	assert.Equal(t, "2024-12-15T16:00:00Z", response.DATA[0].DeliveryDate, "поле не соответствует ожидаемому")
	assert.Equal(t, "abc123xyz", response.DATA[0].OrderKey, "поле не соответствует ожидаемому")
}

func TestGetStatuses(t *testing.T) {
	jsonMock := `
	{
		"DATA": [
			{
			  "group": "OrderStatus",
			  "status_id": 1,
			  "name": "Pending",
			  "description": "The order is awaiting processing."
			},
			{
			  "group": "OrderStatus 2",
			  "status_id": "2",
			  "name": "Pending2",
			  "description": "The order 2 is awaiting processing."
			}
		]
	}`

	ts := NewTestServer("/get_statuses", []byte(jsonMock))
	defer ts.Close()

	response, err := ts.Service.GetStatuses()

	assert.NoError(t, err, "Неожиданная ошибка")
	assert.NotNil(t, response, "Ответ не получен")
	assert.Equal(t, 2, len(response.DATA), "длина массива не совпадает")
	assert.IsTypef(t, responses.GetStatusesResponse{}, *response, "тип ответа не соответствует ожидаемому")
	assert.Equal(t, "OrderStatus", response.DATA[0].Group, "поле не соответствует ожидаемому")
	assert.Equal(t, uint16(1), response.DATA[0].StatusID, "поле не соответствует ожидаемому")
	assert.Equal(t, "Pending", response.DATA[0].Name, "поле не соответствует ожидаемому")
	assert.Equal(t, "The order is awaiting processing.", response.DATA[0].Description, "поле не соответствует ожидаемому")
	assert.Equal(t, "OrderStatus 2", response.DATA[1].Group, "поле не соответствует ожидаемому")
	assert.Equal(t, uint16(2), response.DATA[1].StatusID, "поле не соответствует ожидаемому")
	assert.Equal(t, "Pending2", response.DATA[1].Name, "поле не соответствует ожидаемому")
	assert.Equal(t, "The order 2 is awaiting processing.", response.DATA[1].Description, "поле не соответствует ожидаемому")
}
