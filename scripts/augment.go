//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/database/repo"
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
	videos, err := videoRepo.FindNVideosWithoutComments(context.Background(), 50_000)
	if err != nil {
		log.Println("Failed to find videos without comments:", err)
		return
	}

	youtubeClient := youtube.New()
	videoService := video.New(videoRepo, youtubeClient)

	batchSize := 50

	log.Println("Found", len(videos), "videos without comments, resulting in", len(videos)/batchSize, "batches")
	for i := 0; i < len(videos); i += batchSize {
		log.Println("Processing batch", i/batchSize, "i =", i)
		videoIDs := make([]string, 0)
		videoIDToVideo := make(map[string]db.OpenYoutubeDislikesVideo)
		for j := i; j < i+batchSize; j++ {
			vid := videos[j]
			videoIDToVideo[vid.ID] = vid
			videoIDs = append(videoIDs, vid.ID)
		}

		if videoIDs[0] == "" {
			log.Println("videoIDs[0] is empty, skipping")
			break
		}

		videoResp, err := youtubeClient.GetVideosList(videoIDs)
		if err != nil || len(videoResp.Items) == 0 {
			log.Println("Failed to get videos list:", err)
			continue
		}

		channelIDs := make([]string, 0)
		channelIDToVideoID := make(map[string]string)
		for _, videoItem := range videoResp.Items {
			channelIDs = append(channelIDs, videoItem.Snippet.ChannelId)
			channelIDToVideoID[videoItem.Snippet.ChannelId] = videoItem.Id
		}

		channelResp, err := youtubeClient.GetChannelsList(channelIDs)
		if err != nil {
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
		}
	}
}
