package responses

type SearchBrandsResponse = Response[SearchBrand]

type SearchBrand struct {
	Brand string `json:"brand"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}
