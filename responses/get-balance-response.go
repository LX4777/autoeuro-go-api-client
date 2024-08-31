package responses

type GetBalanceResponse = Response[Balance]

type Balance struct {
	Balance      float64 `json:"balance"`
	Credit       float64 `json:"credit"`
	Ordered      float64 `json:"ordered"`
	Reserved     float64 `json:"reserved"`
	Limit        float64 `json:"limit"`
	PayTomorrow  float64 `json:"pay_tomorrow"`
	ShippingFrom float64 `json:"shipping_from"`
	Active       int     `json:"active"` // 0 | 1
}
