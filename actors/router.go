package actors

import "github.com/go-chi/chi/v5"

func Router() chi.Router {
	router := chi.NewRouter()
	router.Get("/", ListActors)
	router.Get("/{id}", ListActor)
	router.Post("/", CreateActor)

	router.Delete("/{id}", DeleteActor)
	router.Put("/{id}", UpdateActor)

	return router
}
