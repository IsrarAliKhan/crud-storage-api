package requests

import "fmt"

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (c *Category) Validate() error {
	if c.Id == 0 || c.Name == "" {
		return fmt.Errorf("missing required field")
	}
	return nil
}
