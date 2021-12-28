package data

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
)

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views int64, comments *int64, subscribers int64, publishedAt int64) error
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.OpenYoutubeDislikesVideo, error)
}

type MLService interface {
	Predict(ctx context.Context, apiVersion int, video types.Video) (int64, error)
}

type Service struct {
	videoRepo VideoRepo
	mlService MLService
}

func New(mlService MLService, videoRepo VideoRepo) *Service {
	return &Service{
		mlService: mlService,
		videoRepo: videoRepo,
	}
}

func (s *Service) GetDislikes(ctx context.Context, apiVersion int, video types.Video) (int64, string, error) {
	prediction, err := s.mlService.Predict(ctx, apiVersion, video)
	if err != nil {
		return 0, "", err
	}

	predictionString := fmt.Sprintf("%d", prediction)
	if prediction > 1_000_000 {
		predictionString = fmt.Sprintf("%dM", prediction/1_000_000)
	} else if prediction > 1000 {
		predictionString = fmt.Sprintf("%dK", prediction/1000)
	}

	return prediction, predictionString, nil
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
