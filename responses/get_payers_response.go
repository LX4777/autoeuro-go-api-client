package responses

type GetPayersResponse = Response[Payer]

type Payer struct {
	PayerName string `json:"payer_name"`
	PayerKey  string `json:"payer_key"`
}
