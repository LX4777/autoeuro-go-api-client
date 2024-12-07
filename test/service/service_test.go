package service

import (
	"encoding/json"
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
