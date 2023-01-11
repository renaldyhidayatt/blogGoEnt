package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/ent/category"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
)

type CategoryRepository = interfaces.ICategoryRepository

type categoryRepository struct {
	db      *ent.Client
	context context.Context
}

func NewCategoryRepository(db *ent.Client, context context.Context) *categoryRepository {
	return &categoryRepository{db: db, context: context}
}

func (r *categoryRepository) FindAll() ([]*ent.Category, error) {
	category, err := r.db.Category.Query().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result category: %w", err)
	}

	return category, err
}

func (r *categoryRepository) Create(input request.CategoryRequest) (*ent.Category, error) {
	category, err := r.db.Category.Query().Where(category.NameEQ(input.Name)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by name: %w", err)
	}
	if category.ID != 0 {
		return nil, fmt.Errorf("name already exits")
	}

	res, err := r.db.Category.Create().SetName(input.Name).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category create: %w", err)
	}

	return res, nil
}

func (r *categoryRepository) FindById(id int) (*ent.Category, error) {
	category, err := r.db.Category.Query().Where(category.IDEQ(id)).WithPosts().First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}

	return category, nil
}

func (r *categoryRepository) UpdateById(id int, input request.CategoryRequest) (*ent.Category, error) {
	_, err := r.db.Category.Query().Where(category.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}
	user, err := r.db.Category.UpdateOneID(id).SetName(input.Name).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update category: %w", err)
	}

	return user, nil
}

func (r *categoryRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.Category.Query().Where(category.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query category by id: %w", err)
	}

	err = r.db.Category.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete; %w", err)
	}

	return true, nil
}
