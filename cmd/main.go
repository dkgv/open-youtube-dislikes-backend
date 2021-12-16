package main

import (
	"log"

	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/endpoints"
	"github.com/dkgv/dislikes/internal/logic/data"
	"github.com/dkgv/dislikes/internal/logic/ml"
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
	mlService, err := ml.New()
	if err != nil {
		log.Print(err)
	}

	dataService := data.New(mlService, singleDislikeRepo, youtubeVideoRepo)

	// Initialize API
	api := endpoints.New(dataService)
	api.Launch()
}
