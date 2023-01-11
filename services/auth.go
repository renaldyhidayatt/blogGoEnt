package services

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
	"github.com/renaldyhidayatt/blogGoEnt/repository"
)

type AuthService = interfaces.IAuthService

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) RegisterUser(input request.UserRequest) (*ent.User, error) {
	var userRequest request.UserRequest

	userRequest.FirstName = input.FirstName
	userRequest.LastName = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := s.repository.RegisterUser(userRequest)

	return res, err
}

func (s *authService) LoginUser(input request.AuthRequest) (*ent.User, error) {
	var authRequest request.AuthRequest

	authRequest.Email = input.Email
	authRequest.Password = input.Password

	res, err := s.repository.LoginUser(authRequest)

	return res, err
}
