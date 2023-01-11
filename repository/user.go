package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/ent/user"
	"github.com/renaldyhidayatt/blogGoEnt/interfaces"
	"github.com/renaldyhidayatt/blogGoEnt/security"
)

type UserRepository = interfaces.IUserRepository

type userRepository struct {
	db      *ent.Client
	context context.Context
}

func NewUserRepository(db *ent.Client, context context.Context) *userRepository {
	return &userRepository{db: db, context: context}
}

func (r *userRepository) FindAll() ([]*ent.User, error) {
	user, err := r.db.User.Query().WithPosts().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result user: %w", err)
	}

	return user, nil
}

func (r *userRepository) Create(input request.UserRequest) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.EmailEQ(input.Email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by email: %w", err)
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("email already exitst")
	}

	res, err := r.db.User.Create().SetFirstname(input.FirstName).SetLastname(input.LastName).SetEmail(input.Email).SetPassword(security.HashPassword(input.Password)).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user create: %w", err)
	}

	return res, nil
}

func (r *userRepository) FindById(id int) (*ent.User, error) {
	user, err := r.db.User.Query().WithPosts().Where(user.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by id: %w", err)
	}

	return user, nil
}

func (r *userRepository) UpdateById(id int, input request.UserRequest) (*ent.User, error) {
	_, err := r.db.User.Query().Where(user.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by id: %w", err)
	}

	user, err := r.db.User.UpdateOneID(id).SetFirstname(input.FirstName).SetLastname(input.LastName).SetEmail(input.Email).SetPassword(input.Password).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed update user: %w", err)
	}

	return user, nil

}

func (r *userRepository) DeleteById(id int) (bool, error) {
	_, err := r.db.User.Query().Where(user.IDEQ(id)).First(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query user by id: %w", err)
	}

	err = r.db.User.DeleteOneID(id).Exec(r.context)

	if err != nil {
		return false, fmt.Errorf("failed query delete: %w", err)
	}

	return true, nil
}
