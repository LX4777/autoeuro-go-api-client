package requests_data

type CreateOrderRequestData struct {
	DeliveryKey  string      `json:"delivery_key"`
	PayerKey     string      `json:"payer_key"`
	StockItems   []StockItem `json:"stock_items"`
	WaitAllGoods *int        `json:"wait_all_goods,omitempty"` // 0 | 1
	Comment      *string     `json:"comment,omitempty"`
	DeliveryDate *string     `json:"delivery_date,omitempty"`
}

type StockItem struct {
	OfferKey string   `json:"offer_key"`
	Quantity int      `json:"quantity"`
	Price    *float64 `json:"price,omitempty"`
	Comment  *string  `json:"comment,omitempty"`
}
