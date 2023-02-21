package films

import (
	"errors"
	"net/http"

	"example.com/gosql/db"
	"example.com/gosql/myerrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func ListFilm(w http.ResponseWriter, r *http.Request) {
	var film *Film
	id := chi.URLParam(r, "id")
	db.DB.First(&film, id)

	if film.FilmId == 0 {
		render.Render(w, r, myerrors.ErrInvalidRequest(errors.New("Film with the specified id doesn't exist")))
	} else {
		render.Render(w, r, NewFilmResponse(film))
	}
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, myerrors.ErrInvalidRequest(err))
	}

	film := data.Film

	db.DB.Create(film)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmResponse(film))
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DB.Delete(&Film{}, id)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	DeleteFilm(w, r)
	CreateFilm(w, r)
}
