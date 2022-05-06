package dislikes

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/database"
	"github.com/dkgv/dislikes/internal/logic/ml"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/dkgv/dislikes/internal/types"
)

type VideoRepo interface {
	FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.OpenYoutubeDislikesVideo, error)
	FindByID(ctx context.Context, id string) (db.OpenYoutubeDislikesVideo, error)
}

type MLService interface {
	Predict(ctx context.Context, apiVersion ml.ModelType, video types.Video) (int64, error)
}

type DislikeRepo interface {
	GetDislikeCount(ctx context.Context, videoID string) (int64, error)
}

type CommentRepo interface {
	FindCommentStatusByVideoID(ctx context.Context, videoID string) (bool, error)
	FindSentimentByVideoID(ctx context.Context, videoID string) (db.FindSentimentByVideoIDRow, error)
}

type Service struct {
	videoRepo   VideoRepo
	mlService   MLService
	dislikeRepo DislikeRepo
	commentRepo CommentRepo
}

func New(mlService MLService, videoRepo VideoRepo, dislikeRepo DislikeRepo, commentRepo CommentRepo) *Service {
	return &Service{
		mlService:   mlService,
		videoRepo:   videoRepo,
		dislikeRepo: dislikeRepo,
		commentRepo: commentRepo,
	}
}

func (s *Service) GetDislikes(ctx context.Context, videoID string) (int64, string, error) {
	exactDislikes, err := s.retrieveExactDislikes(ctx, videoID)
	if err == nil {
		return exactDislikes, formatDislikes(exactDislikes), nil
	}

	return s.retrieveEstimatedDislikes(ctx, videoID)
}

func (s *Service) retrieveEstimatedDislikes(ctx context.Context, videoID string) (int64, string, error) {
	dbVideo, err := s.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return 0, "0", err
	}

	exists, err := s.commentRepo.FindCommentStatusByVideoID(ctx, videoID)
	if err != nil {
		return 0, "0", err
	}

	video := mappers.DBVideoToVideo(dbVideo)

	var modelType ml.ModelType
	if exists {
		modelType = ml.ModelTypeSentiment

		row, err := s.commentRepo.FindSentimentByVideoID(ctx, videoID)
		if err != nil {
			return 0, "0", err
		}

		video.Positive, _ = strconv.ParseFloat(row.Positive, 64)
		video.Negative, _ = strconv.ParseFloat(row.Negative, 64)
		video.Neutral, _ = strconv.ParseFloat(row.Neutral, 64)
		video.Compound, _ = strconv.ParseFloat(row.Compound, 64)
	} else {
		modelType = ml.ModelTypeSimple
	}

	predictedDislikes, err := s.mlService.Predict(ctx, modelType, video)
	if err != nil {
		return 0, "0", err
	}

	if video.Likes+predictedDislikes > video.Views {
		predictedDislikes = 0
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

func (s *Service) GetDislikeEstimationsByHash(ctx context.Context, apiVersion ml.ModelType, video types.Video, count int32) ([]types.DislikeEstimation, error) {
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
