package films

import (
	"errors"
	"fmt"
	"net/http"

	"example.com/gosql/db"
	"example.com/gosql/myerrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	var films []*Film
	db.DB.Where("title LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&films)
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
	//Read the data from the request
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, myerrors.ErrInvalidRequest(err))
	}
	film := data.Film
	id := film.FilmId

	var existingFilm Film
	db.DB.First(&existingFilm, id)
	if existingFilm.FilmId != 0 {
		render.Status(r, http.StatusForbidden)
		render.Render(w, r, myerrors.ErrInvalidRequest(errors.New("Film with the specified id already axists")))
		return
	}

	//Will assign a FilmId if it is 0
	db.DB.Create(film)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmResponse(film))
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DB.Delete(&Film{}, id)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	//Read the data from the request
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, myerrors.ErrInvalidRequest(err))
	}
	film := data.Film
	id := chi.URLParam(r, "id")

	var existingFilm Film
	db.DB.First(&existingFilm, id)
	if existingFilm.FilmId == 0 {
		render.Status(r, http.StatusForbidden)
		render.Render(w, r, myerrors.ErrInvalidRequest(errors.New("cannot update nonexistent movie")))
		return
	}
	existingFilm = *film
	db.DB.Save(&existingFilm)

	render.Status(r, http.StatusAccepted)
	render.Render(w, r, NewFilmResponse(film))
}
