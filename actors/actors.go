package actors

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type Actor struct {
	ActorId            int    `gorm:"type:smallint;primaryKey"`
	Title              string `gorm:"type:varchar(128)"`
	LanguageId         int    `gorm:"type:tinyint"`
	Description        string `gorm:"type:longtext"`
	ReleaseYear        string `gorm:"type:longtext"`
	OriginalLanguageId int    `gorm:"type:tinyint"`
	RentalDuration     int
	RentalRate         float64
	Length             int
	ReplacementCost    float64 `gorm:"type:double"`
	Rating             string
	SpecialFeatures    string
	LastUpdate         string
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
