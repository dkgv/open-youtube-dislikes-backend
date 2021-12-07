package main

import (
	"github.com/dkgv/dislikes/internal/api"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	singleDislikeRepo := repo.NewSingleDislikeRepo(conn)
	aggregateDislikeRepo := repo.NewAggregateDislikeRepo(conn)
	youtubeVideoRepo := repo.NewYouTubeVideoRepo(conn)

	api := api.NewAPI(singleDislikeRepo, aggregateDislikeRepo, youtubeVideoRepo)
	err = api.Start()
	if err != nil {
		panic(err)
	}
}
