package response

import "github.com/renaldyhidayatt/blogGoEnt/ent"

type CategoryResponse struct {
	Name string `json:"name"`
	Post []Post
}

type Post struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Image string `json:"img"`
	Body  string `json:"body"`
}

func ResponseCategory(category *ent.Category) CategoryResponse {

	for _, post := range category.Edges.Posts {
		return CategoryResponse{
			Name: category.Name,
			Post: []Post{
				{Title: post.Title, Slug: post.Slug, Image: post.Img, Body: post.Body},
			},
		}
	}
	return CategoryResponse{}
}
