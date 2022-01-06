package dislikes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
	"github.com/dkgv/dislikes/internal/youtube"
)

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views int64, comments int64, subscribers int64, publishedAt int64) error
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.OpenYoutubeDislikesVideo, error)
	FindByID(ctx context.Context, id string) (db.OpenYoutubeDislikesVideo, error)
}

type MLService interface {
	Predict(ctx context.Context, apiVersion int, video types.Video) (int64, error)
}

type DislikeRepo interface {
	GetDislikeCount(ctx context.Context, videoID string) (int64, error)
}

type YouTubeClient interface {
	GetStatistics(videoID string) (*youtube.StatisticsResponse, error)
}

type Service struct {
	videoRepo     VideoRepo
	mlService     MLService
	dislikeRepo   DislikeRepo
	youtubeClient YouTubeClient
}

func New(mlService MLService, videoRepo VideoRepo, dislikeRepo DislikeRepo, youtubeClient YouTubeClient) *Service {
	return &Service{
		mlService:     mlService,
		videoRepo:     videoRepo,
		dislikeRepo:   dislikeRepo,
		youtubeClient: youtubeClient,
	}
}

func (s *Service) GetDislikes(ctx context.Context, apiVersion int, videoID string) (int64, string, error) {
	exactDislikes, err := s.retrieveExactDislikes(ctx, videoID)
	if err == nil {
		return exactDislikes, formatDislikes(exactDislikes), nil
	}

	return s.retrieveEstimatedDislikes(ctx, apiVersion, videoID)
}

func (s *Service) retrieveEstimatedDislikes(ctx context.Context, apiVersion int, videoID string) (int64, string, error) {
	dbVideo, err := s.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return 0, "0", err
	}

	video := mappers.DBVideoToVideo(dbVideo)
	predictedDislikes, err := s.mlService.Predict(ctx, apiVersion, video)
	if err != nil {
		return 0, "0", err
	}

	return predictedDislikes, "~" + formatDislikes(predictedDislikes), nil
}

func (s *Service) retrieveExactDislikes(ctx context.Context, videoID string) (int64, error) {
	dbVideo, err := s.videoRepo.FindByID(ctx, videoID)
	if database.IsNoRowError(err) {
		return 0, err
	}

	historicDislikes := dbVideo.Dislikes
	if historicDislikes == 0 {
		return 0, errors.New("no historic dislikes")
	}

	extensionDislikes, err := s.dislikeRepo.GetDislikeCount(ctx, videoID)
	if err != nil {
		return historicDislikes, nil
	}

	return historicDislikes + extensionDislikes, nil
}

func formatDislikes(dislikes int64) string {
	formated := fmt.Sprintf("%d", dislikes)
	if dislikes > 1_000_000 {
		formated = fmt.Sprintf("%dM", dislikes/1_000_000)
	} else if dislikes > 1000 {
		formated = fmt.Sprintf("%dK", dislikes/1000)
	}
	return formated
}

func (s *Service) GetDislikeEstimationsByHash(ctx context.Context, apiVersion int, video types.Video, count int32) ([]types.DislikeEstimation, error) {
	// Retrieve at most 5 videos matching hash
	count = int32(math.Min(5.0, float64(count)))
	dbVideos, err := s.videoRepo.FindNByHash(ctx, video.IDHash, count)
	if err != nil {
		return nil, err
	}

	estimations := make([]types.DislikeEstimation, len(dbVideos))
	for i := range dbVideos {
		dislikes, err := s.mlService.Predict(ctx, apiVersion, mappers.DBVideoToVideo(dbVideos[i]))
		if err != nil {
			continue
		}

		estimations[i] = types.DislikeEstimation{
			IDHash:   dbVideos[i].IDHash,
			Dislikes: dislikes,
		}
	}
	return estimations, nil
}

func (s *Service) AddVideo(ctx context.Context, videoID string, video types.Video) error {
	if video.Comments < 0 {
		err := s.augmentVideo(videoID, &video)
		if err != nil {
			log.Printf("Failed to augment video %s: %s", videoID, err)
		}
	}

	videoIDHash := hashString(videoID)
	return s.videoRepo.Upsert(
		ctx,
		videoID,
		videoIDHash,
		video.Likes,
		video.Dislikes,
		video.Views,
		video.Comments,
		video.Subscribers,
		video.PublishedAt,
	)
}

func (s *Service) augmentVideo(videoID string, video *types.Video) error {
	resp, err := s.youtubeClient.GetStatistics(videoID)
	if err != nil {
		return err
	}

	if len(resp.Items) == 0 {
		return errors.New("no statistics found")
	}

	statistics := resp.Items[0].Statistics
	commentCountString := statistics.CommentCount
	commentCount, err := strconv.ParseInt(commentCountString, 10, 64)
	if err != nil {
		commentCount = 0
	}

	viewCountString := statistics.ViewCount
	viewCount, err := strconv.ParseInt(viewCountString, 10, 64)
	if err != nil {
		viewCount = 0
	}

	video.Comments = commentCount
	video.Views = viewCount
	return nil
}

func hashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
