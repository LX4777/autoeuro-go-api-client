package responses

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type SearchItemsResponse = Response[SearchItem]

type SearchItem struct {
	OfferKey        string  `json:"offer_key"`
	Stock           int     `json:"stock"` // 0 | 1
	Cross           *uint8  `json:"cross"` // null | 0 | 1 | 2 | 3 | 10 | 11 | 12, по факту приходит string
	Brand           string  `json:"brand"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	Packing         int     `json:"packing"`
	Price           float32 `json:"price"` // по факту приходит string
	Currency        string  `json:"currency"`
	Amount          int     `json:"amount"`
	Unit            string  `json:"unit"`
	Return          uint8   `json:"return"` // 0 | 1,  по факту приходит string
	OrderBefore     string  `json:"order_before"`
	DeliveryTime    string  `json:"delivery_time"`
	DeliveryTimeMax string  `json:"delivery_time_max"`
	Rejects         float32 `json:"rejects"`
	Dealer          int     `json:"dealer"` // 0 | 1
	WarehouseName   string  `json:"warehouse_name"`
	WarehouseKey    string  `json:"warehouse_key"`
}

func (s *SearchItem) UnmarshalJSON(data []byte) error {
	type Alias SearchItem
	aux := &struct {
		Cross  interface{} `json:"cross"`
		Price  string      `json:"price"`
		Return string      `json:"return"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Обработка поля Cross
	if aux.Cross != nil {
		switch v := aux.Cross.(type) {
		case string:
			if v == "null" {
				s.Cross = nil
			} else {
				val, err := strconv.ParseUint(v, 10, 8)
				if err != nil {
					return fmt.Errorf("failed to convert cross to uint8: %w", err)
				}
				cross := uint8(val)
				s.Cross = &cross
			}
		case float64:
			cross := uint8(v)
			s.Cross = &cross
		}
	}

	// Обработка поля Price
	price, err := strconv.ParseFloat(aux.Price, 32)
	if err != nil {
		return fmt.Errorf("failed to convert price to float32: %w", err)
	}
	s.Price = float32(price)

	// Обработка поля Return
	ret, err := strconv.ParseUint(aux.Return, 10, 8)
	if err != nil {
		return fmt.Errorf("failed to convert return to uint8: %w", err)
	}
	s.Return = uint8(ret)

	return nil
}
