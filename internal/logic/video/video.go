package video

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"unicode"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/dkgv/dislikes/internal/youtube"
	"github.com/grassmudhorses/vader-go/sentitext"
)

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views int64, comments int64, subscribers int64, publishedAt int64, durationSec int32) error
}

type CommentRepo interface {
	Insert(ctx context.Context, videoID string, content string, negative float32, neutral float32, positive float32, compound float32) error
}

type YouTubeClient interface {
	GetVideosList(videoIDs []string) (*youtube.VideosListResponse, error)
	GetChannelsList(channelIDs []string) (*youtube.ChannelsListResponse, error)
	GetCommentThreadForVideo(videoID string, count int) (*youtube.CommentThreadResponse, error)
}

type MLService interface {
	Sentiment(ctx context.Context, text string) (sentitext.Sentiment, error)
}

type Service struct {
	videoRepo     VideoRepo
	commentRepo   CommentRepo
	youtubeClient YouTubeClient
	mlService     MLService
}

func New(videoRepo VideoRepo, youtubeClient YouTubeClient, mlService MLService, commentRepo CommentRepo) *Service {
	return &Service{
		videoRepo:     videoRepo,
		youtubeClient: youtubeClient,
		mlService:     mlService,
		commentRepo:   commentRepo,
	}
}

func (s *Service) ProcessVideo(ctx context.Context, videoID string, video types.Video) error {
	if video.Views == 0 || video.Comments <= 0 {
		err := s.AugmentVideo(videoID, &video)
		if err != nil {
			log.Printf("Failed to augment video %s: %s", videoID, err)
			return err
		}
	}

	err := s.ProcessVideoComments(ctx, videoID)
	if err != nil {
		log.Printf("Failed to process video %s comments: %s", videoID, err)
		return err
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
		video.DurationSec,
	)
}

func (s *Service) ProcessVideoComments(ctx context.Context, videoID string) error {
	resp, err := s.youtubeClient.GetCommentThreadForVideo(videoID, 99)
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}

	for _, comment := range resp.Comments {
		content := comment.Snippet.TopLevelComment.Snippet.TextOriginal

		// Sentiment analysis only works with English comments
		if containsRussian(content) {
			continue
		}

		sentiment, err := s.mlService.Sentiment(context.Background(), content)
		if err != nil {
			continue
		}

		// Discard error for single comment insertion
		_ = s.commentRepo.Insert(context.Background(),
			videoID,
			content,
			float32(sentiment.Negative),
			float32(sentiment.Neutral),
			float32(sentiment.Positive),
			float32(sentiment.Compound),
		)
	}

	return nil
}

func hashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func containsRussian(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}
