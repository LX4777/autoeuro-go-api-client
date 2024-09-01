package responses

type GetOrdersResponse = Response[OrderDetail]

type OrderDetail struct {
	Brand        string  `json:"brand"`
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Amount       int     `json:"amount"`
	Unit         string  `json:"unit"`
	Dealer       int     `json:"dealer"`     // 0 | 1
	Cancelable   int     `json:"cancelable"` // 0 | 1
	Returnable   int     `json:"returnable"` // 0 | 1
	StatusID     int     `json:"status_id"`
	Status       string  `json:"status"`
	Document     string  `json:"document"`
	OrderID      int     `json:"order_id"`
	Comment      string  `json:"comment"`
	United       int     `json:"united"` // 0 | 1
	OrderDate    string  `json:"order_date"`
	OrderNumber  string  `json:"order_number"`
	Delivery     string  `json:"delivery"`
	DeliveryDate string  `json:"delivery_date"`
	OrderKey     string  `json:"order_key"`
}
