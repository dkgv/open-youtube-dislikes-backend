package data

import (
	"context"
	"crypto/sha256"

	"github.com/dkgv/dislikes/internal/types"
)

type SingleDislikeRepo interface {
	Insert(ctx context.Context, id string, hashedIP string) error
}

type YouTubeVideoRepo interface {
	Upsert(ctx context.Context, id string, likes, dislikes, views, comments int64) error
}

type MLService interface {
	Predict(ctx context.Context, details types.VideoDetails) (int64, error)
}

type Service struct {
	singleDislikeRepo SingleDislikeRepo
	youTubeVideoRepo  YouTubeVideoRepo
	mlService         MLService
}

func New(mlService MLService, singleDislikeRepo SingleDislikeRepo, youTubeVideoRepo YouTubeVideoRepo) *Service {
	return &Service{
		mlService:         mlService,
		singleDislikeRepo: singleDislikeRepo,
		youTubeVideoRepo:  youTubeVideoRepo,
	}
}

func (s *Service) GetDislikes(ctx context.Context, details types.VideoDetails) (types.VideoDetails, error) {
	prediction, err := s.mlService.Predict(ctx, details)
	if err != nil {
		return types.VideoDetails{}, err
	}

	return types.VideoDetails{
		Dislikes: prediction,
	}, nil
}

func (s *Service) AddDislike(ctx context.Context, videoID string, ip string) error {
	hashedIPBytes := sha256.Sum256([]byte(ip))
	hashedIP := string(hashedIPBytes[:])
	return s.singleDislikeRepo.Insert(ctx, videoID, hashedIP)
}

func (s *Service) AddYouTubeVideo(ctx context.Context, videoID string, details types.VideoDetails) error {
	return s.youTubeVideoRepo.Upsert(
		ctx,
		videoID,
		details.Likes,
		details.Dislikes,
		details.Views,
		details.Comments,
	)
}
