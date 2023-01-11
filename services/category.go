package services

import (
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
	"github.com/renaldyhidayatt/blogGoEnt/repository"
)

type CategoryService = interfaces.ICategoryService

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *categoryService {
	return &categoryService{repository: repository}
}

func (s *categoryService) FindAll() ([]*ent.Category, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *categoryService) FindById(id int) (*ent.Category, error) {
	res, err := s.repository.FindById(id)

	return res, err
}

func (s *categoryService) Create(input request.CategoryRequest) (*ent.Category, error) {
	var categoryRequest request.CategoryRequest

	categoryRequest.Name = input.Name

	res, err := s.repository.Create(categoryRequest)

	return res, err
}

func (s *categoryService) UpdateById(id int, input request.CategoryRequest) (*ent.Category, error) {
	var categoryRequest request.CategoryRequest

	categoryRequest.Name = input.Name

	res, err := s.repository.UpdateById(id, categoryRequest)

	return res, err
}

func (s *categoryService) DeleteById(id int) (bool, error) {
	res, err := s.repository.DeleteById(id)

	return res, err
}
