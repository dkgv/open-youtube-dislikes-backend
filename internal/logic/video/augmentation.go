package video

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/dkgv/dislikes/internal/youtube"
)

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
