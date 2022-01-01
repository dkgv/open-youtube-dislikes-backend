package dislikes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
)

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views int64, comments *int64, subscribers int64, publishedAt int64) error
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.OpenYoutubeDislikesVideo, error)
	FindByID(ctx context.Context, id string) (db.OpenYoutubeDislikesVideo, error)
}

type MLService interface {
	Predict(ctx context.Context, apiVersion int, video types.Video) (int64, error)
}

type DislikeRepo interface {
	GetDislikeCount(ctx context.Context, videoID string) (int64, error)
}

type Service struct {
	videoRepo   VideoRepo
	mlService   MLService
	dislikeRepo DislikeRepo
}

func New(mlService MLService, videoRepo VideoRepo, dislikeRepo DislikeRepo) *Service {
	return &Service{
		mlService:   mlService,
		videoRepo:   videoRepo,
		dislikeRepo: dislikeRepo,
	}
}

func (s *Service) GetDislikes(ctx context.Context, apiVersion int, videoID string, video types.Video) (int64, string, error) {
	exactDislikes, err := s.retrieveExactAmount(ctx, videoID)
	if err == nil {
		return exactDislikes, formatDislikes(exactDislikes), nil
	}

	predictedDislikes, err := s.mlService.Predict(ctx, apiVersion, video)
	if err == nil {
		return predictedDislikes, "~" + formatDislikes(predictedDislikes), nil
	}

	return 0, "0", err
}

func (s *Service) retrieveExactAmount(ctx context.Context, videoID string) (int64, error) {
	dbVideo, err := s.videoRepo.FindByID(ctx, videoID)
	if database.IsNoRowError(err) {
		return 0, err
	}

	historicDislikes := dbVideo.Dislikes
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

func (s *Service) AddVideo(ctx context.Context, videoID string, details types.Video) error {
	videoIDHash := hashString(videoID)
	return s.videoRepo.Upsert(
		ctx,
		videoID,
		videoIDHash,
		details.Likes,
		details.Dislikes,
		details.Views,
		details.Comments,
		details.Subscribers,
		details.PublishedAt,
	)
}

func hashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
