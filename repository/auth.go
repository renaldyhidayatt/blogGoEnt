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

type AuthRepository = interfaces.IAuthRepository

type authRepository struct {
	db      *ent.Client
	context context.Context
}

func NewAuthRepository(db *ent.Client, context context.Context) *authRepository {
	return &authRepository{db: db, context: context}
}

func (r *authRepository) RegisterUser(input request.UserRequest) (*ent.User, error) {

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

func (r *authRepository) LoginUser(input request.AuthRequest) (*ent.User, error) {
	user, err := r.db.User.Query().Where(user.EmailEQ(input.Email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user by email: %w", err)
	}

	checkPassword := security.VerifyPassword(user.Password, input.Password)

	if checkPassword != nil {
		return nil, fmt.Errorf("failed checkhash password: %w", err)

	}

	return user, nil
}
