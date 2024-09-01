package responses

type GetBrandsResponse = Response[Brand]

type Brand struct {
	Brand string `json:"brand"`
}
