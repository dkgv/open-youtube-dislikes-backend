package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

func NewAPI() *API {
	return &API{
		router: mux.NewRouter(),
	}
}

func (a *API) Start() error {
	a.router.HandleFunc("/dislike", a.AddDislike).Methods("POST")
	a.router.HandleFunc("/dislikes", a.GetDislikes).Methods("GET")

	err := http.ListenAndServe(":9000", a.router)
	return err
}
