package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func (c *Client) GetStatistics(videoID string) (*StatisticsResponse, error) {
	url := c.buildURL(videoID)
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

	statistics := &StatisticsResponse{}
	err = json.NewDecoder(resp.Body).Decode(statistics)
	if err != nil {
		return nil, err
	}

	return statistics, nil
}

func (c *Client) buildURL(videoID string) string {
	url := "https://youtube.googleapis.com/youtube/v3/videos?part=statistics&id=%s&key=%s"
	return fmt.Sprintf(url, videoID, c.key)
}
