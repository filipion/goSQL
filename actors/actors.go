package actors

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type Actor struct {
	ActorId    int    `gorm:"type:smallint;primaryKey"`
	FirstName  string `gorm:"type:varchar(45)"`
	LastName   string `gorm:"type:varchar(45)"`
	LastUpdate string
}

func (Actor) TableName() string {
	return "actor"
}

type ActorRequest struct {
	*Actor
}

func (f *ActorRequest) Bind(r *http.Request) error {
	if f.Actor == nil {
		return errors.New("missing required Actor fields")
	}
	return nil
}

type ActorResponse struct {
	*Actor
}

func NewActorResponse(actor *Actor) *ActorResponse {
	return &ActorResponse{actor}
}

func NewActorListResponse(actors []*Actor) []render.Renderer {
	list := []render.Renderer{}
	for _, actor := range actors {
		list = append(list, NewActorResponse(actor))
	}
	return list
}

func (f *ActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
