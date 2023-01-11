package interfaces

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
)

type IPostRepository interface {
	FindAll() ([]*ent.Post, error)
	FindById(id int) (*ent.Post, error)
	Create(input request.PostRequest) (*ent.Post, error)
	UpdateById(id int, input request.PostRequest) (*ent.Post, error)
	DeleteById(id int) (bool, error)
}

type IPostService interface {
	FindAll() ([]*ent.Post, error)
	FindById(id int) (*ent.Post, error)
	Create(input request.PostRequest) (*ent.Post, error)
	UpdateById(id int, input request.PostRequest) (*ent.Post, error)
	DeleteById(id int) (bool, error)
}
