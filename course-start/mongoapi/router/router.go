package router

import (
	"mongo/controller"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/api/movies", controller.GetMyAllMovies)
	router.Post("/api/movie", controller.CreateMovie)
	router.Put("/api/movie/{id}", controller.MarkedAsWatched)
	router.Delete("/api/movie/{id}", controller.DeleteAmovie)
	router.Delete("/api/deleteallmovie", controller.DeleteAllmovie)

	return router

}
