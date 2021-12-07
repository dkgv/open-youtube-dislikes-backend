package api

import (
	"net/http"

	"github.com/dkgv/dislikes/internal/database"
	"github.com/gorilla/mux"
)

type API struct {
	router               *mux.Router
	dislikeRepo          database.DislikeRepo
	aggregateDislikeRepo database.AggregateDislikeRepo
}

func NewAPI(dislikeRepo database.DislikeRepo, aggregateDislikeRepo database.AggregateDislikeRepo) *API {
	return &API{
		router:               mux.NewRouter(),
		dislikeRepo:          dislikeRepo,
		aggregateDislikeRepo: aggregateDislikeRepo,
	}
}

func (a *API) Start() error {
	a.router.HandleFunc("/dislike", a.AddDislike).Methods("POST")
	a.router.HandleFunc("/dislikes", a.GetDislikes).Methods("GET")

	err := http.ListenAndServe(":9000", a.router)
	return err
}
