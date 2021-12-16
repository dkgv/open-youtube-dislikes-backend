package endpoints

import (
	"log"
	"net/http"

	"github.com/dkgv/dislikes/internal/logic/data"
	"github.com/gorilla/mux"
)

type API struct {
	dataService *data.Service
}

func New(statsService *data.Service) *API {
	return &API{
		dataService: statsService,
	}
}

func (a *API) Launch() {
	routes := a.defineRoutes()
	log.Fatal(http.ListenAndServe(":5000", routes))
}

func (a *API) defineRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})
	router.HandleFunc("/video/{id}/add", a.PostVideoAdd).Methods("POST")
	router.HandleFunc("/video/{id}/add_dislike", a.PostVideoAddDislike).Methods("POST")
	router.HandleFunc("/video/{id}/estimate_dislikes", a.GetVideoEstimateDislikes).Methods("GET")
	return router
}
