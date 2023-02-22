package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId             int    `gorm:"type:smallint;primaryKey"`
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
	LastUpdate         time.Time `gorm:"autoUpdateTime"`
	Actors             []*Actor  `gorm:"many2many:film_actor"`
}

func (Film) TableName() string {
	return "film"
}

type FilmRequest struct {
	*Film
}

func (f *FilmRequest) Bind(r *http.Request) error {
	if f.Film == nil {
		return errors.New("missing required Film fields")
	}
	return nil
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(film *Film) *FilmResponse {
	return &FilmResponse{film}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
