package responses

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type GetStatusesResponse = Response[Status]

type Status struct {
	Group       string `json:"group"`
	StatusID    uint16 `json:"status_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Status) UnmarshalJSON(data []byte) error {
	type Alias Status
	aux := &struct {
		StatusID any `json:"status_id"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	switch v := aux.StatusID.(type) {
	case string:
		val, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return fmt.Errorf("ошибка конвертации поля status_id в uint16: %w", err)
		}
		s.StatusID = uint16(val)
	case float64:
		s.StatusID = uint16(v)
	default:
		return fmt.Errorf("ошибка: неожиданный для поля status_id тип: %T", v)
	}

	return nil
}
