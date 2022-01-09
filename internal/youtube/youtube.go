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
	url := c.buildURL("videos", videoIDs)
	rawResp, err := baseRequest(url)
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	resp := &VideosListResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetChannelsList(channelIDs []string) (*ChannelsListResponse, error) {
	url := c.buildURL("channels", channelIDs)
	rawResp, err := baseRequest(url)
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	resp := &ChannelsListResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func baseRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func (c *Client) buildURL(endpoint string, ids []string) string {
	idsParam := strings.Join(ids, ",")
	maxResults := int(math.Min(float64(len(ids)), 50))
	url := "https://youtube.googleapis.com/youtube/v3/%s?part=statistics,contentDetails,snippet&id=%s&key=%s&maxResults=%d"
	return fmt.Sprintf(url, endpoint, idsParam, c.key, maxResults)
}
