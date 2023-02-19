package requests

import "fmt"

type Item struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

func (i *Item) Validate() error {
	if i.Id == 0 || i.Name == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}
