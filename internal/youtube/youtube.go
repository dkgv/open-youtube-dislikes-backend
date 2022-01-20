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
	key    string
	client *http.Client
}

func New() *Client {
	key := os.Getenv("YOUTUBE_API_KEY")
	return &Client{
		key:    key,
		client: &http.Client{},
	}
}

func (c *Client) GetVideosList(videoIDs []string) (*VideosListResponse, error) {
	url := c.buildBulkURL("videos", []string{"statistics,contentDetails,snippet"}, videoIDs)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	rawResp, err := c.client.Do(req)
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
	url := c.buildBulkURL("channels", []string{"statistics"}, channelIDs)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	rawResp, err := c.client.Do(req)
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

func (c *Client) GetCommentThreadForVideo(videoID string, count int) (*CommentThreadResponse, error) {
	parameters := fmt.Sprintf("videoId=%s&maxResults=%d&textFormat=plainText&part=snippet&order=relevance", videoID, count)
	url := c.buildURL("commentThreads", parameters)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	rawResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	resp := &CommentThreadResponse{}
	err = json.NewDecoder(rawResp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CanFind(videoID string) (bool, error) {
	url := fmt.Sprintf("https://i1.ytimg.com/vi/%s/hqdefault.jpg", videoID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Accept", "application/json")

	rawResp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}

	return rawResp.StatusCode == 200, nil
}

func (c *Client) buildBulkURL(endpoint string, parts []string, ids []string) string {
	idsParam := strings.Join(ids, ",")
	partsParam := strings.Join(parts, ",")
	maxResults := int(math.Min(float64(len(ids)), 50))
	parameters := fmt.Sprintf("part=%s&id=%s&maxResults=%d", partsParam, idsParam, maxResults)
	return c.buildURL(endpoint, parameters)
}

func (c *Client) buildURL(endpoint string, parameters string) string {
	url := "https://youtube.googleapis.com/youtube/v3/%s?key=%s&%s"
	return fmt.Sprintf(url, endpoint, c.key, parameters)
}
