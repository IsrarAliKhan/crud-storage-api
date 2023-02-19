package responses

type Article struct {
	Id       uint64  `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}
