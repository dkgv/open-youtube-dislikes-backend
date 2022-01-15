package video

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/dkgv/dislikes/internal/youtube"
)

type VideoRepo interface {
	Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views int64, comments int64, subscribers int64, publishedAt int64, durationSec int32) error
}

type YouTubeClient interface {
	GetVideosList(videoIDs []string) (*youtube.VideosListResponse, error)
	GetChannelsList(channelIDs []string) (*youtube.ChannelsListResponse, error)
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
		err := s.AugmentVideo(videoID, &video)
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
		video.DurationSec,
	)
}

func hashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func (s *Service) AugmentVideo(videoID string, video *types.Video) error {
	videoResp, err := s.youtubeClient.GetVideosList([]string{videoID})
	if err != nil {
		return err
	}

	if len(videoResp.Items) == 0 {
		return fmt.Errorf("video %s not found", videoID)
	}

	videoItem := videoResp.Items[0]

	channelResp, err := s.youtubeClient.GetChannelsList([]string{videoItem.Snippet.ChannelId})
	if err != nil {
		return err
	}
	channelItem := channelResp.Items[0]

	return s.AugmentVideoStruct(videoItem, channelItem, video)
}

func (s *Service) AugmentVideoStruct(videoItem youtube.VideoItem, channelItem youtube.ChannelItem, video *types.Video) error {
	statistics := videoItem.Statistics
	contentDetails := videoItem.ContentDetails

	likeCount := parseInt64(statistics.LikeCount)
	viewCount := parseInt64(statistics.ViewCount)
	commentCount := parseInt64(statistics.CommentCount)
	subscribers := parseInt64(channelItem.Statistics.SubscriberCount)
	publishedAt := parseDateToMillis(videoItem.Snippet.PublishedAt)
	durationSec := parseDurationToSec(contentDetails.Duration)

	video.Likes = max(likeCount, video.Likes)
	video.Views = max(viewCount, video.Views)
	video.Comments = max(commentCount, video.Comments)
	video.Subscribers = max(subscribers, video.Subscribers)
	video.PublishedAt = max(publishedAt, video.PublishedAt)
	video.DurationSec = int32(max(durationSec, int64(video.DurationSec)))

	return nil
}

func parseDateToMillis(date time.Time) int64 {
	return date.UnixNano() / int64(time.Millisecond)
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
