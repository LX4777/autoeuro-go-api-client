package responses

type Meta struct {
	Resource   string                 `json:"resource"`
	Version    string                 `json:"version"`
	Section    string                 `json:"section"`
	Parameters map[string]interface{} `json:"parameters"`
	Date       string                 `json:"date"`
	UserID     string                 `json:"user_id"`
}

type Response[T any] struct {
	META  Meta `json:"META"`
	DATA  []T  `json:"DATA,omitempty"`
	ERROR *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"ERROR,omitempty"`
}
