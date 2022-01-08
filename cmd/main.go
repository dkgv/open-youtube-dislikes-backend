package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/logic/dislikes"
	"github.com/dkgv/dislikes/internal/logic/ml"
	"github.com/dkgv/dislikes/internal/logic/user"
	"github.com/dkgv/dislikes/internal/logic/video"
	"github.com/dkgv/dislikes/internal/swagger"
	"github.com/dkgv/dislikes/internal/youtube"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			return
		}
	}

	conn, err := database.NewConnection()
	if err != nil {
		log.Println(err)
		return
	}

	dislikeRepo := repo.NewDislikeRepo(conn)
	likeRepo := repo.NewLikeRepo(conn)
	videoRepo := repo.NewVideoRepo(conn)
	userRepo := repo.NewUserRepo(conn)

	mlService, err := ml.New()
	if err != nil {
		return
	}

	youtubeClient := youtube.New()
	dislikeService := dislikes.New(mlService, videoRepo, dislikeRepo)
	userService := user.New(userRepo, likeRepo, dislikeRepo)
	videoService := video.New(videoRepo, youtubeClient)

	instance := swagger.New(dislikeService, userService, videoService)
	instance.Run()

	log.Println("Server started successfully")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	sig := <-signals
	log.Printf("Got signal: %s", sig)
}
