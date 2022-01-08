package main

import (
	"context"

	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/logic/video"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/youtube"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		return
	}

	videoRepo := repo.NewVideoRepo(conn)
	videos, err := videoRepo.FindNVideosWithoutComments(context.Background(), 50_000)
	if err != nil {
		return
	}

	youtubeClient := youtube.New()
	videoService := video.New(videoRepo, youtubeClient)

	batchSize := 50
	for i := 0; i < len(videos); i += batchSize {
		batchIDs := make([]string, batchSize)
		for j := i; j < batchSize; j++ {
			batchIDs[j-i] = videos[j].ID
		}

		resp, err := youtubeClient.GetVideosList(batchIDs)
		if err != nil {
			continue
		}

		if len(resp.Items) != batchSize {
			continue
		}

		batchVideos := videos[i : i+batchSize]
		for j := range batchIDs {
			dbVideo := batchVideos[j]
			vid := mappers.DBVideoToVideo(dbVideo)
			err = videoService.AugmentVideoStruct(dbVideo.ID, &vid)
			if err != nil {
				continue
			}

			_ = videoRepo.Upsert(context.Background(),
				dbVideo.ID,
				vid.IDHash,
				vid.Likes,
				vid.Dislikes,
				vid.Views,
				vid.Comments,
				vid.Subscribers,
				vid.PublishedAt,
			)
		}
	}
}
