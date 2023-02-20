package films

import "github.com/go-chi/chi/v5"

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)
	router.Get("/{id}", ListFilm)
	router.Post("/", CreateFilm)

	return router
}
