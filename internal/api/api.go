package api

import (
	"net/http"

	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/gorilla/mux"
)

type API struct {
	dislikeRepo          *repo.SingleDislikeRepo
	aggregateDislikeRepo *repo.AggregateDislikeRepo
	youtubeDislikeRepo   *repo.YouTubeVideoRepo
}

func New(dislikeRepo *repo.SingleDislikeRepo, aggregateDislikeRepo *repo.AggregateDislikeRepo, youtubeDislikeRepo *repo.YouTubeVideoRepo) *API {
	return &API{
		dislikeRepo:          dislikeRepo,
		aggregateDislikeRepo: aggregateDislikeRepo,
		youtubeDislikeRepo:   youtubeDislikeRepo,
	}
}

func (a *API) DefineRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("Hello World"))
	})
	router.HandleFunc("/add_single_dislike", a.AddSingleDislike).Methods("POST")
	router.HandleFunc("/add_youtube_video", a.AddYouTubeVideo).Methods("POST")
	return router
}
