package responses

type GetStatusesResponse = Response[Status]

type Status struct {
	Group       string `json:"group"`
	StatusID    int    `json:"status_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
