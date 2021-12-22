package main

import (
	"log"

	"github.com/dkgv/dislikes/internal/api"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
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
	videoRepo := repo.NewVideoRepo(conn)

	// Define services
	mlService, err := ml.New()
	if err != nil {
		log.Print(err)
	}

	dataService := data.New(mlService, singleDislikeRepo, videoRepo)

	instance := api.New(dataService)
	instance.Run()
}
