package interfaces

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
)

type ICategoryRepository interface {
	FindAll() ([]*ent.Category, error)
	FindById(id int) (*ent.Category, error)
	Create(input request.CategoryRequest) (*ent.Category, error)
	UpdateById(id int, input request.CategoryRequest) (*ent.Category, error)
	DeleteById(id int) (bool, error)
}

type ICategoryService interface {
	FindAll() ([]*ent.Category, error)
	FindById(id int) (*ent.Category, error)
	Create(input request.CategoryRequest) (*ent.Category, error)
	UpdateById(id int, input request.CategoryRequest) (*ent.Category, error)
	DeleteById(id int) (bool, error)
}
