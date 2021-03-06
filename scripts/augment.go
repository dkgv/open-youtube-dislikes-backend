//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"
	"unicode"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/logic/ml"
	"github.com/dkgv/dislikes/internal/logic/video"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/youtube"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return
	}

	videoRepo := repo.NewVideoRepo(conn)
	commentRepo := repo.NewCommentRepo(conn)

	videos, err := videoRepo.FindNVideosMissingDataWithDislikes(context.Background(), 50_000)
	if err != nil {
		log.Println("Failed to find videos without comments:", err)
		return
	}

	youtubeClient := youtube.New()
	mlService, err := ml.New()
	videoService := video.New(videoRepo, youtubeClient, mlService, commentRepo)
	if err != nil {
		log.Println("Failed to create ml service:", err)
		return
	}

	batchSize := 40
	for i := 0; i < len(videos)-batchSize; i += batchSize {
		log.Println("Processing batch", i/batchSize, "i =", i)

		videoIDs := make([]string, 0)
		videoIDToVideo := make(map[string]db.OpenYoutubeDislikesVideo)
		index := i
		for index < i+batchSize {
			vid := videos[index]
			index++

			exists, err := youtubeClient.CanFind(vid.ID)
			if !exists || err != nil {
				_ = videoRepo.DeleteVideoByID(context.Background(), vid.ID)
				continue
			}

			videoIDs = append(videoIDs, vid.ID)
			videoIDToVideo[vid.ID] = vid
		}

		if len(videoIDs) == 0 {
			log.Println("No videos to process, next batch")
			continue
		}

		log.Println("VideoIDs:", videoIDs)
		videoResp, err := youtubeClient.GetVideosList(videoIDs)
		if err != nil || len(videoResp.Items) == 0 {
			log.Println("Failed to get videos list:", err)
			continue
		}

		channelIDs := make([]string, 0)
		channelIDToVideoID := make(map[string]string)
		for _, videoItem := range videoResp.Items {
			channelID := videoItem.Snippet.ChannelId
			if _, ok := channelIDToVideoID[channelID]; !ok {
				channelIDs = append(channelIDs, channelID)
				channelIDToVideoID[channelID] = videoItem.Id
			}
		}

		channelResp, err := youtubeClient.GetChannelsList(channelIDs)
		if err != nil {
			log.Println("Failed to get channels list:", err)
			continue
		}

		videoIDToChannelItem := make(map[string]youtube.ChannelItem)
		for _, channelItem := range channelResp.Items {
			videoIDToChannelItem[channelIDToVideoID[channelItem.Id]] = channelItem
		}

		for _, videoItem := range videoResp.Items {
			channelItem, ok := videoIDToChannelItem[videoItem.Id]
			if !ok {
				continue
			}

			dbVideo := videoIDToVideo[videoItem.Id]
			vid := mappers.DBVideoToVideo(dbVideo)
			err = videoService.AugmentVideoStruct(videoItem, channelItem, &vid)
			if err != nil {
				log.Println("Failed to augment video:", err)
				continue
			}

			log.Println("Augmented video:", dbVideo.ID)
			_ = videoRepo.Upsert(context.Background(),
				dbVideo.ID,
				vid.IDHash,
				vid.Likes,
				vid.Dislikes,
				vid.Views,
				vid.Comments,
				vid.Subscribers,
				vid.PublishedAt,
				vid.DurationSec,
			)

			_ = videoService.ProcessVideoComments(context.Background(), dbVideo.ID)
		}
	}
}

func containsRussian(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}
