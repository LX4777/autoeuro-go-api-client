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
		Cross   any `json:"cross"`
		Price   any `json:"price"`
		Return  any `json:"return"`
		Packing any `json:"packing"`
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
					return fmt.Errorf("ошибка конвертации поля Cross из string в uint8: %w", err)
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
	switch v := aux.Price.(type) {
	case string:
		price, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return fmt.Errorf("ошибка конвертации поля Price из string в float32: %w", err)
		}
		s.Price = float32(price)
	case float32:
		s.Price = v
	case float64:
		s.Price = float32(v)
	default:
		return fmt.Errorf("ошибка конвертации поля Price: неизвестный тип %T", v)
	}

	// Обработка поля Return
	switch v := aux.Return.(type) {
	case string:
		ret, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return fmt.Errorf("ошибка конвертации поля Return в uint8: %w", err)
		}
		s.Return = uint8(ret)
	case uint8:
		s.Return = v
	case int:
		s.Return = uint8(v)
	case float64:
		s.Return = uint8(v)
	default:
		return fmt.Errorf("ошибка конвертации поля Return: неизвестный тип %T", v)
	}

	// Обработка поля Packing
	switch v := aux.Packing.(type) {
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("ошибка конвертации поля Packing из string в int64: %w (значение поля: %v)", err, v)
		}
		s.Packing = int(i)
	case int:
		s.Packing = v
	case float64:
		s.Packing = int(v)
	default:
		return fmt.Errorf("ошибка конвертации поля Packing: неизвестный тип %T", v)
	}

	return nil
}
