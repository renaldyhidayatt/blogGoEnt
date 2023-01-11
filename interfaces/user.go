package interfaces

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
)

type IUserRepository interface {
	FindAll() ([]*ent.User, error)
	Create(input request.UserRequest) (*ent.User, error)
	FindById(id int) (*ent.User, error)
	UpdateById(id int, input request.UserRequest) (*ent.User, error)
	DeleteById(id int) (bool, error)
}

type IUserService interface {
	FindAll() ([]*ent.User, error)
	Create(input request.UserRequest) (*ent.User, error)
	FindById(id int) (*ent.User, error)
	UpdateById(id int, input request.UserRequest) (*ent.User, error)
	DeleteById(id int) (bool, error)
}
