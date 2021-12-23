package data

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/database/repo"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
)

type SingleDislikeRepo interface {
	Insert(ctx context.Context, videoID string, userID string) error
	Delete(ctx context.Context, videoID string, userID string) error
}

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views, comments, subscribers, publishedAt uint32) error
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.Video, error)
}

type UserRepo interface {
	FindByID(ctx context.Context, id string) (db.User, error)
	Insert(ctx context.Context, id string) error
}

type MLService interface {
	Predict(ctx context.Context, apiVersion int, video types.Video) (uint32, error)
}

type Service struct {
	singleDislikeRepo SingleDislikeRepo
	videoRepo         VideoRepo
	mlService         MLService
	userRepo          UserRepo
}

func New(mlService MLService, singleDislikeRepo SingleDislikeRepo, videoRepo VideoRepo, userRepo *repo.UserRepo) *Service {
	return &Service{
		mlService:         mlService,
		singleDislikeRepo: singleDislikeRepo,
		videoRepo:         videoRepo,
		userRepo:          userRepo,
	}
}

func (s *Service) GetDislikes(ctx context.Context, apiVersion int, video types.Video) (string, error) {
	prediction, err := s.mlService.Predict(ctx, apiVersion, video)
	if err != nil {
		return "", err
	}

	predictionString := fmt.Sprintf("%d", prediction)
	if prediction > 1000000 {
		prediction /= 1000000
		predictionString = fmt.Sprintf("%dM", prediction)
	} else if prediction > 1000 {
		prediction /= 1000
		predictionString = fmt.Sprintf("%dK", prediction)
	}

	return predictionString, nil
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

func (s *Service) AddDislike(ctx context.Context, videoID string, userID string) error {
	err := s.userRepo.Insert(ctx, userID)
	if err != nil {
		return err
	}

	return s.singleDislikeRepo.Insert(ctx, videoID, userID)
}

func (s *Service) RemoveDislike(ctx context.Context, videoID string, userID string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil || user == (db.User{}) {
		return err
	}

	return s.singleDislikeRepo.Delete(ctx, videoID, userID)
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
	hashedInputBytes := sha256.Sum256([]byte(input))
	return string(hashedInputBytes[:])
}
