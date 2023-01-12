package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/ent/category"
	"github.com/renaldyhidayatt/blogGoEnt/ent/post"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
)

type PostRepository = interfaces.IPostRepository

type postsRepository struct {
	db      *ent.Client
	context context.Context
}

func NewPostRepository(db *ent.Client, context context.Context) *postsRepository {
	return &postsRepository{db: db, context: context}
}

func (r *postsRepository) FindAll() ([]*ent.Post, error) {
	post, err := r.db.Post.Query().WithCategories().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result post: %w", err)
	}

	return post, err
}

func (r *postsRepository) Create(input request.PostRequest) (*ent.Post, error) {
	category, err := r.db.Category.Query().Where(category.IDEQ(input.CategoryId)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}

	res, err := r.db.Post.Create().SetTitle(input.Title).SetSlug(input.Slug).SetImg(input.Image).SetBody(input.Body).AddCategories(category).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query post create: %w", err)
	}

	return res, nil
}

func (r *postsRepository) FindById(id int) (*ent.Post, error) {
	res, err := r.db.Post.Query().WithCategories().Where(post.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query post by id: %w", err)
	}

	return res, err
}

func (r *postsRepository) UpdateById(id int, input request.PostRequest) (*ent.Post, error) {
	category, err := r.db.Category.Query().Where(category.IDEQ(input.CategoryId)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query category by id: %w", err)
	}

	_, err = r.db.Post.Query().WithCategories().Where(post.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query post by id: %w", err)
	}
	res, err := r.db.Post.UpdateOneID(id).SetTitle(input.Title).SetSlug(input.Slug).AddCategories(category).SetImg(input.Image).SetBody(input.Body).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update post: %w", err)
	}

	return res, err
}

func (r *postsRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.Post.Query().WithCategories().Where(post.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query post by id: %w", err)
	}

	err = r.db.Post.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete: %w", err)
	}

	return true, nil
}
