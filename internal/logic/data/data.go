package data

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
)

type DislikeRepo interface {
	Insert(ctx context.Context, videoID string, userID string) error
	Delete(ctx context.Context, videoID string, userID string) error
	FindByID(ctx context.Context, videoID string, userID string) (db.OpenYoutubeDislikesDislike, error)
}

type LikeRepo interface {
	Insert(ctx context.Context, videoID string, userID string) error
	Delete(ctx context.Context, videoID string, userID string) error
	FindByID(ctx context.Context, videoID string, userID string) (db.OpenYoutubeDislikesLike, error)
}

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views, comments, subscribers, publishedAt uint32) error
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.OpenYoutubeDislikesVideo, error)
}

type UserRepo interface {
	FindByID(ctx context.Context, id string) (db.OpenYoutubeDislikesUser, error)
	Insert(ctx context.Context, id string) error
}

type MLService interface {
	Predict(ctx context.Context, apiVersion int, video types.Video) (uint32, error)
}

type Service struct {
	dislikeRepo DislikeRepo
	likeRepo    LikeRepo
	videoRepo   VideoRepo
	mlService   MLService
	userRepo    UserRepo
}

func New(mlService MLService, dislikeRepo DislikeRepo, likeRepo LikeRepo, videoRepo VideoRepo, userRepo UserRepo) *Service {
	return &Service{
		mlService:   mlService,
		dislikeRepo: dislikeRepo,
		likeRepo:    likeRepo,
		videoRepo:   videoRepo,
		userRepo:    userRepo,
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

	return s.dislikeRepo.Insert(ctx, videoID, userID)
}

func (s *Service) RemoveDislike(ctx context.Context, videoID string, userID string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil || user == (db.OpenYoutubeDislikesUser{}) {
		return err
	}

	return s.dislikeRepo.Delete(ctx, videoID, userID)
}

func (s *Service) AddLike(ctx context.Context, videoID string, userID string) error {
	err := s.userRepo.Insert(ctx, userID)
	if err != nil {
		return err
	}

	return s.likeRepo.Insert(ctx, videoID, userID)
}

func (s *Service) RemoveLike(ctx context.Context, videoID string, userID string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil || user == (db.OpenYoutubeDislikesUser{}) {
		return err
	}

	return s.likeRepo.Delete(ctx, videoID, userID)
}

func (s *Service) HasDislikedVideo(ctx context.Context, videoID string, userID string) (bool, error) {
	dislike, err := s.dislikeRepo.FindByID(ctx, videoID, userID)
	if err != nil {
		return false, err
	}

	return dislike != (db.OpenYoutubeDislikesDislike{}), nil
}

func (s *Service) HasLikedVideo(ctx context.Context, videoID string, userID string) (bool, error) {
	dislike, err := s.likeRepo.FindByID(ctx, videoID, userID)
	if err != nil {
		return false, err
	}

	return dislike != (db.OpenYoutubeDislikesLike{}), nil
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
