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
	dislikeRepo := repo.NewDislikeRepo(conn)
	likeRepo := repo.NewLikeRepo(conn)
	videoRepo := repo.NewVideoRepo(conn)
	userRepo := repo.NewUserRepo(conn)

	// Define services
	mlService, err := ml.New()
	if err != nil {
		log.Print(err)
	}

	dataService := data.New(mlService, dislikeRepo, likeRepo, videoRepo, userRepo)

	instance := api.New(dataService)
	instance.Run()
}
