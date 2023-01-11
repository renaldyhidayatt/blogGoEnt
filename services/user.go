package services

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
	"github.com/renaldyhidayatt/blogGoEnt/repository"
)

type UserService = interfaces.IUserService

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) FindAll() ([]*ent.User, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *userService) FindById(id int) (*ent.User, error) {
	res, err := s.repository.FindById(id)

	return res, err
}

func (s *userService) Create(input request.UserRequest) (*ent.User, error) {
	var userRequest request.UserRequest

	userRequest.FirstName = input.FirstName
	userRequest.LastName = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := s.repository.Create(userRequest)

	return res, err
}

func (s *userService) UpdateById(id int, input request.UserRequest) (*ent.User, error) {
	var userRequest request.UserRequest

	userRequest.FirstName = input.FirstName
	userRequest.LastName = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := s.repository.UpdateById(id, userRequest)

	return res, err
}

func (s *userService) DeleteById(id int) (bool, error) {

	res, err := s.repository.DeleteById(id)

	return res, err
}
