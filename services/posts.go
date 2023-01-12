package services

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
	"github.com/renaldyhidayatt/blogGoEnt/repository"
)

type PostService = interfaces.IPostService

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *postService {
	return &postService{repository: repository}
}

func (s *postService) FindAll() ([]*ent.Post, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *postService) FindById(id int) (*ent.Post, error) {
	res, err := s.repository.FindById(id)

	return res, err
}

func (s *postService) Create(input request.PostRequest) (*ent.Post, error) {
	var postRequest request.PostRequest

	postRequest.Title = input.Title
	postRequest.Slug = input.Slug
	postRequest.Image = input.Image
	postRequest.Body = input.Body
	postRequest.CategoryId = input.CategoryId

	res, err := s.repository.Create(postRequest)

	return res, err
}

func (s *postService) UpdateById(id int, input request.PostRequest) (*ent.Post, error) {
	var postRequest request.PostRequest

	postRequest.Title = input.Title
	postRequest.Slug = input.Slug
	postRequest.Image = input.Image
	postRequest.Body = input.Body
	postRequest.CategoryId = input.CategoryId

	res, err := s.repository.FindById(id)

	return res, err
}

func (s *postService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
