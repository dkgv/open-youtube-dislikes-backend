package main

import (
	"log"

	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/endpoints"
	"github.com/dkgv/dislikes/internal/logic/data"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		log.Print(err)
	}

	// Define repositories
	singleDislikeRepo := repo.NewSingleDislikeRepo(conn)
	youtubeVideoRepo := repo.NewYouTubeVideoRepo(conn)

	// Define services
	dataService := data.New(singleDislikeRepo, youtubeVideoRepo)

	// Initialize API
	api := endpoints.New(dataService)
	api.Launch()
}
