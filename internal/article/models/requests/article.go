package requests

import "fmt"

type Article struct {
	Id       uint64  `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}

func (a *Article) Validate() error {
	if a.Id == 0 || a.Name == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}
