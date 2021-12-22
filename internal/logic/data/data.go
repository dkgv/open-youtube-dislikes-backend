package data

import (
	"context"
	"crypto/sha256"
	"math"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
)

type SingleDislikeRepo interface {
	Insert(ctx context.Context, id string, hashedIP string) error
	Delete(ctx context.Context, id string, hashedIP string) error
}

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views, comments, subscribers uint32) error
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.Video, error)
}

type MLService interface {
	Predict(ctx context.Context, apiVersion int, video types.Video) (uint32, error)
}

type Service struct {
	singleDislikeRepo SingleDislikeRepo
	videoRepo         VideoRepo
	mlService         MLService
}

func New(mlService MLService, singleDislikeRepo SingleDislikeRepo, videoRepo VideoRepo) *Service {
	return &Service{
		mlService:         mlService,
		singleDislikeRepo: singleDislikeRepo,
		videoRepo:         videoRepo,
	}
}

func (s *Service) PredictDislikes(ctx context.Context, apiVersion int, video types.Video) (uint32, error) {
	prediction, err := s.mlService.Predict(ctx, apiVersion, video)
	if err != nil {
		return 0, err
	}

	return prediction, nil
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
		dislikes, err := s.PredictDislikes(ctx, apiVersion, mappers.DBVideoToVideo(dbVideos[i]))
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

func (s *Service) AddDislike(ctx context.Context, videoID string, ip string) error {
	return s.singleDislikeRepo.Insert(ctx, videoID, hashString(ip))
}

func (s *Service) RemoveDislike(ctx context.Context, videoID string, ip string) error {
	return s.singleDislikeRepo.Delete(ctx, videoID, hashString(ip))
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
	)
}

func hashString(input string) string {
	// TODO: don't mask IP by hashing, can be bruteforced
	hashedInputBytes := sha256.Sum256([]byte(input))
	return string(hashedInputBytes[:])
}
