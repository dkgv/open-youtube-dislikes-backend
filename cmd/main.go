package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/logic/data"
	"github.com/dkgv/dislikes/internal/logic/ml"
	"github.com/dkgv/dislikes/internal/logic/user"
	"github.com/dkgv/dislikes/internal/swagger"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		log.Println(err)
		return
	}

	// Define repositories
	dislikeRepo := repo.NewDislikeRepo(conn)
	likeRepo := repo.NewLikeRepo(conn)
	videoRepo := repo.NewVideoRepo(conn)
	userRepo := repo.NewUserRepo(conn)

	// Define services
	mlService, err := ml.New()
	if err != nil {
		return
	}

	dataService := data.New(mlService, videoRepo)
	userService := user.New(userRepo, likeRepo, dislikeRepo)

	instance := swagger.New(dataService, userService)
	instance.Run()

	log.Println("Server started successfully")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	sig := <-signals
	log.Printf("Got signal: %s", sig)
}
