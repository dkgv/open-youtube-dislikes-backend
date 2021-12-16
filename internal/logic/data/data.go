package data

import (
	"context"

	"github.com/dkgv/dislikes/internal/types"
)

type SingleDislikeRepo interface {
	AddDislike(ctx context.Context, id string, hashedIP string) error
}

type YouTubeVideoRepo interface {
	AddYouTubeVideo(ctx context.Context, id string, likes, dislikes, views, comments int64) error
}

type Service struct {
	singleDislikeRepo SingleDislikeRepo
	youTubeVideoRepo  YouTubeVideoRepo
}

func New(singleDislikeRepo SingleDislikeRepo, youTubeVideoRepo YouTubeVideoRepo) *Service {
	return &Service{
		singleDislikeRepo: singleDislikeRepo,
		youTubeVideoRepo:  youTubeVideoRepo,
	}
}

func (s *Service) EstimateDislikes(ctx context.Context, details types.VideoDetails) (types.VideoDetails, error) {
	dislikes := int64(0)

	return types.VideoDetails{
		Dislikes: dislikes,
	}, nil
}

func (s *Service) AddDislike(ctx context.Context, videoID string) error {
	return s.singleDislikeRepo.AddDislike(ctx, videoID, "")
}

func (s *Service) AddYouTubeVideo(ctx context.Context, videoID string, details types.VideoDetails) error {
	return s.youTubeVideoRepo.AddYouTubeVideo(
		ctx,
		videoID,
		details.Likes,
		details.Dislikes,
		details.Views,
		details.Comments,
	)
}
