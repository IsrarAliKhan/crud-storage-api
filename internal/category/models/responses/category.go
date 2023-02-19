package responses

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
