package request

import "errors"

type PostRequest struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Image      string `json:"img"`
	Body       string `json:"body"`
	CategoryId int    `json:"category_id"`
}

func (p *PostRequest) Validate() error {
	if p.Title == "" {
		return errors.New("title tidak boleh kosong")
	} else if p.Slug == "" {
		return errors.New("slug tidak boleh kosong")
	} else if p.Image == "" {
		return errors.New("image tidak boleh kosong")
	} else if p.Body == "" {
		return errors.New("body tidak boleh kosong")
	}

	return nil
}
