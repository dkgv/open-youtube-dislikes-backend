package api

import (
	"net/http"

	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/gorilla/mux"
)

type API struct {
	router               *mux.Router
	dislikeRepo          *repo.SingleDislikeRepo
	aggregateDislikeRepo *repo.AggregateDislikeRepo
	youtubeDislikeRepo   *repo.YouTubeVideoRepo
}

func NewAPI(dislikeRepo *repo.SingleDislikeRepo, aggregateDislikeRepo *repo.AggregateDislikeRepo, youtubeDislikeRepo *repo.YouTubeVideoRepo) *API {
	return &API{
		router:               mux.NewRouter(),
		dislikeRepo:          dislikeRepo,
		aggregateDislikeRepo: aggregateDislikeRepo,
		youtubeDislikeRepo:   youtubeDislikeRepo,
	}
}

func (a *API) Start() error {
	a.router.HandleFunc("/dislike", a.AddSingleDislike).Methods("POST")
	a.router.HandleFunc("/add_youtube_video", a.AddYouTubeVideo).Methods("POST")

	err := http.ListenAndServe(":9000", a.router)
	return err
}
