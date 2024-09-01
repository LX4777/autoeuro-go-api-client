package responses

type CreateOrderResponse = Response[Order]

type Order struct {
	OrderID           int    `json:"order_id"`
	Result            bool   `json:"result"`
	ResultDescription string `json:"result_description"`
}
