package requests

type GetOrdersRequestData struct {
	Orders  []int64        `json:"orders,omitempty"`
	Filters []OrderFilters `json:"filters,omitempty"`
}

type OrderFilters struct {
	From        string  `json:"from"`
	To          string  `json:"to"`
	DeliveryKey string  `json:"delivery_key"`
	PayerKey    *string `json:"payer_key,omitempty"`
}
