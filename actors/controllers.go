package actors

import (
	"errors"
	"net/http"

	"example.com/gosql/db"
	"example.com/gosql/myerrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	render.RenderList(w, r, NewActorListResponse(actors))
}

func ListActor(w http.ResponseWriter, r *http.Request) {
	var actor *Actor
	id := chi.URLParam(r, "id")
	db.DB.First(&actor, id)

	if actor.ActorId == 0 {
		render.Render(w, r, myerrors.ErrInvalidRequest(errors.New("Actor with the specified id doesn't exist")))
	} else {
		render.Render(w, r, NewActorResponse(actor))
	}
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, myerrors.ErrInvalidRequest(err))
	}
	actor := data.Actor
	id := actor.ActorId

	var existingActor Actor
	db.DB.First(&existingActor, id)
	if existingActor.ActorId != 0 {
		render.Render(w, r, myerrors.ErrInvalidRequest(errors.New("Actor with the specified id already axists")))
		return
	}

	db.DB.Create(actor)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActorResponse(actor))
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DB.Delete(&Actor{}, id)
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	//Read the data from the request
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, myerrors.ErrInvalidRequest(err))
	}
	actor := data.Actor
	id := chi.URLParam(r, "id")

	var existingActor Actor
	db.DB.First(&existingActor, id)
	if existingActor.ActorId == 0 {
		render.Status(r, http.StatusForbidden)
		render.Render(w, r, myerrors.ErrInvalidRequest(errors.New("cannot update nonexistent movie")))
		return
	}

	existingActor = *actor
	db.DB.Save([]Actor{existingActor})
	render.Status(r, http.StatusAccepted)
	render.Render(w, r, NewActorResponse(actor))
}
