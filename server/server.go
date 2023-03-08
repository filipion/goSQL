package server

import (
	"log"
	"net/http"
	"os"

	"example.com/gosql/models"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Mount("/films", models.FilmsRouter())
	router.Mount("/actors", models.ActorsRouter())

	return router
}

func Init() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	router := Router()

	log.Printf("Server starting at http://10.0.0.4:%s\n", port)

	err := http.ListenAndServe("0.0.0.0:"+port, router)

	log.Fatal(err)
}
