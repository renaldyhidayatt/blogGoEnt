package routes

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/handler"
	"github.com/renaldyhidayatt/blogGoEnt/repository"
	"github.com/renaldyhidayatt/blogGoEnt/services"
)

func NewAuthRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewAuthRepository(db, context)

	service := services.NewAuthService(repository)
	handler := handler.NewAuthHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Post("/login", handler.LoginUser)
		r.Post("/register", handler.RegisterUser)
	})
}