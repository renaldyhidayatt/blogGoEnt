package routes

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/blogGoEnt/ent"
	"github.com/renaldyhidayatt/blogGoEnt/handler"
	"github.com/renaldyhidayatt/blogGoEnt/middlewares"
	"github.com/renaldyhidayatt/blogGoEnt/repository"
	"github.com/renaldyhidayatt/blogGoEnt/services"
)

func NewCategoryRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewCategoryRepository(db, context)
	service := services.NewCategoryService(repository)
	handler := handler.NewCategoryHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", handler.FindAll)
		r.Get("/{id:[0-9]+}", handler.FindById)
		r.Post("/", handler.Create)
		r.Put("/{id:[0-9]+}", handler.UpdateById)
		r.Delete("/{id:[0-9]+}", handler.DeleteById)
	})
}
