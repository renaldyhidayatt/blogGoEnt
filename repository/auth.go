package repository

import (
	"context"

	"entgo.io/ent/examples/privacytenant/ent"
)

type authRepository struct {
	db      *ent.Client
	context context.Context
}

func NewAuthRepository(db *ent.Client, context context.Context) *authRepository {
	return &authRepository{db: db, context: context}
}

func (r *authRepository) RegisterUser() (*ent.User, error) {
	var userModel ent.User

}
