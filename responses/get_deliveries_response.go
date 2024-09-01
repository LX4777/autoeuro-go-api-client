package responses

type GetDeliveriesResponse = Response[Delivery]

type Delivery struct {
	DeliveryKey  string `json:"delivery_key"`
	Name         string `json:"name"`
	TimeShiftMSK int    `json:"time_shift_msk"`
}
