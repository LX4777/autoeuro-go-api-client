package responses

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type GetDeliveriesResponse = Response[Delivery]

type Delivery struct {
	DeliveryKey  string `json:"delivery_key"`
	Name         string `json:"name"`
	TimeShiftMSK uint8  `json:"time_shift_msk"` //Должно быть int, по факту приходит str
}

func (d *Delivery) UnmarshalJSON(data []byte) error {
	type Alias Delivery
	aux := &struct {
		TimeShiftMsk interface{} `json:"time_shift_msk"` // Используем interface{} для обработки разных типов данных
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Обработка TimeShiftMsk в зависимости от типа
	switch v := aux.TimeShiftMsk.(type) {
	case string:
		shift, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("failed to convert time_shift_msk to int: %w", err)
		}
		d.TimeShiftMSK = uint8(shift)
	case float64:
		d.TimeShiftMSK = uint8(v)
	default:
		return fmt.Errorf("unexpected type for time_shift_msk: %T", v)
	}

	return nil
}
