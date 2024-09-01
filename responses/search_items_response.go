package responses

type SearchItemsResponse = Response[SearchItem]

type SearchItem struct {
	OfferKey        string `json:"offer_key"`
	Stock           int    `json:"stock"` // 0 | 1
	Cross           *int   `json:"cross"` // null | 0 | 1 | 2 | 3 | 10 | 11 | 12
	Brand           string `json:"brand"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	Packing         int    `json:"packing"`
	Price           string `json:"price"`
	Currency        string `json:"currency"`
	Amount          int    `json:"amount"`
	Unit            string `json:"unit"`
	Return          string `json:"return"` // '0' | '1'
	OrderBefore     string `json:"order_before"`
	DeliveryTime    string `json:"delivery_time"`
	DeliveryTimeMax string `json:"delivery_time_max"`
	Rejects         int    `json:"rejects"`
	Dealer          int    `json:"dealer"` // 0 | 1
	WarehouseName   string `json:"warehouse_name"`
	WarehouseKey    string `json:"warehouse_key"`
}
