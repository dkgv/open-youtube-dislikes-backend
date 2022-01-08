package video

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"regexp"
	"strconv"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/dkgv/dislikes/internal/youtube"
)

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views int64, comments int64, subscribers int64, publishedAt int64) error
}

type YouTubeClient interface {
	GetVideosList(videoIDs []string) (*youtube.VideosListResponse, error)
}

type Service struct {
	videoRepo     VideoRepo
	youtubeClient YouTubeClient
}

func New(videoRepo VideoRepo, youtubeClient YouTubeClient) *Service {
	return &Service{
		videoRepo:     videoRepo,
		youtubeClient: youtubeClient,
	}
}

func (s *Service) AddVideo(ctx context.Context, videoID string, video types.Video) error {
	if video.Views == 0 || video.Comments <= 0 {
		err := s.AugmentVideoStruct(videoID, &video)
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

func hashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func (s *Service) AugmentVideoStruct(videoID string, video *types.Video) error {
	resp, err := s.youtubeClient.GetVideosList([]string{videoID})
	if err != nil {
		return err
	}

	if len(resp.Items) == 0 {
		return errors.New("no statistics found")
	}

	item := resp.Items[0]
	statistics := item.Statistics
	contentDetails := item.ContentDetails

	commentCount := parseInt64(statistics.CommentCount)
	viewCount := parseInt64(statistics.ViewCount)
	likeCount := parseInt64(statistics.LikeCount)
	durationSec := parseDurationToSec(contentDetails.Duration)

	video.Comments = max(commentCount, video.Comments)
	video.Views = max(viewCount, video.Views)
	video.Subscribers = max(likeCount, video.Likes)
	video.DurationSec = max(durationSec, video.DurationSec)

	return nil
}

func parseDurationToSec(duration string) int64 {
	re := regexp.MustCompile(`^P(?:(\d+)D)?T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+(?:.\d+)?)S)?$`)
	matches := re.FindStringSubmatch(duration)
	if matches == nil {
		return 0
	}

	seconds := int64(0)
	if matches[1] != "" {
		days := parseInt64(matches[1])
		seconds += days * 24 * 60 * 60
	}

	if matches[2] != "" {
		hours := parseInt64(matches[2])
		seconds += hours * 60 * 60
	}

	if matches[3] != "" {
		minutes := matches[3]
		seconds += parseInt64(minutes) * 60
	}

	if matches[4] != "" {
		seconds += parseInt64(matches[4])
	}

	return seconds
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func parseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return i
}
