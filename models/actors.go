package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Actor struct {
	ActorId    int       `gorm:"type:smallint;primaryKey"`
	FirstName  string    `gorm:"type:varchar(45)"`
	LastName   string    `gorm:"type:varchar(45)"`
	LastUpdate time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	Films      []*Film   `gorm:"many2many:film_actor"`
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
