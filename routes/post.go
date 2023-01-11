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

func NewPostRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewPostRepository(db, context)
	service := services.NewPostService(repository)
	handler := handler.NewPostHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", handler.FindAll)
		r.Get("/:id", handler.FindById)
		r.Post("/", handler.Create)
		r.Put("/:id", handler.UpdateById)
		r.Delete("/:id", handler.DeleteById)
	})
}
