package main

import (
	"log"
	"net/http"

	"github.com/dkgv/dislikes/internal/api"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		log.Print(err)
	}

	log.Print("Established database connection")

	singleDislikeRepo := repo.NewSingleDislikeRepo(conn)
	aggregateDislikeRepo := repo.NewAggregateDislikeRepo(conn)
	youtubeVideoRepo := repo.NewYouTubeVideoRepo(conn)

	log.Print("Created repositories")

	api := api.New(singleDislikeRepo, aggregateDislikeRepo, youtubeVideoRepo)
	routes := api.DefineRoutes()
	log.Fatal(http.ListenAndServe(":5000", routes))
}
