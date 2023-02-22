package models

import "github.com/go-chi/chi/v5"

func FilmsRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/", ListFilms)
	router.Get("/{id}", ListFilm)
	router.Post("/", CreateFilm)

	router.Delete("/{id}", DeleteFilm)
	router.Put("/{id}", UpdateFilm)

	return router
}

func ActorsRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/", ListActors)
	router.Get("/{id}", ListActor)
	router.Post("/", CreateActor)

	router.Delete("/{id}", DeleteActor)
	router.Put("/{id}", UpdateActor)

	return router
}
