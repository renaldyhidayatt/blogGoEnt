package interfaces

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
)

type IAuthRepository interface {
	RegisterUser(input request.UserRequest) (*ent.User, error)
	LoginUser(input request.AuthRequest) (*ent.User, error)
}

type IAuthService interface {
	RegisterUser(input request.UserRequest) (*ent.User, error)
	LoginUser(input request.AuthRequest) (*ent.User, error)
}
