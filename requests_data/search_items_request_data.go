package requests_data

type SearchItemsRequestData struct {
	Brand       string `json:"brand"`
	Code        string `json:"code"`
	DeliveryKey string `json:"delivery_key"`
	WithCrosses *int   `json:"with_crosses,omitempty"` // 0 | 1
	WithOffers  *int   `json:"with_offers,omitempty"`  // 0 | 1
}
