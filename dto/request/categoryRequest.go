package request

import "errors"

type CategoryRequest struct {
	Name string `json:"name"`
}

func (c *CategoryRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name tidak boleh kosong")
	}

	return nil
}
