package response

import (
	"strings"

	"github.com/renaldyhidayatt/blogGoEnt/ent"
)

type PostResponse struct {
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Image    string `json:"img"`
	Body     string `json:"body"`
	Category []Category
}

type Category struct {
	Name string `json:"name"`
}

func ResponsePost(res *ent.Post) PostResponse {
	var categoryNames []string
	for _, category := range res.Edges.Categories {
		categoryNames = append(categoryNames, category.Name)
	}
	responsePost := PostResponse{
		Title: res.Title,
		Slug:  res.Slug,
		Image: res.Img,
		Body:  res.Body,
		Category: []Category{
			{Name: strings.Join(categoryNames, ",")},
		},
	}

	return responsePost
}
