package youtube

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	key string
}

func New() *Client {
	key := os.Getenv("YOUTUBE_API_KEY")
	return &Client{
		key: key,
	}
}

func (c *Client) GetVideosList(videoIDs []string) (*VideosListResponse, error) {
	url := c.buildURL(videoIDs)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statistics := &VideosListResponse{}
	err = json.NewDecoder(resp.Body).Decode(statistics)
	if err != nil {
		return nil, err
	}

	return statistics, nil
}

func (c *Client) buildURL(videoIDs []string) string {
	videoIDsParam := strings.Join(videoIDs, ",")
	maxResults := int(math.Min(float64(len(videoIDs)), 50))
	url := "https://youtube.googleapis.com/youtube/v3/videos?part=statistics,contentDetails&id=%s&key=%s&maxResults=%d"
	return fmt.Sprintf(url, videoIDsParam, c.key, maxResults)
}
