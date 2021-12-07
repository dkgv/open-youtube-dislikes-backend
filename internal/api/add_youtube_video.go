package api

import (
	"context"
	"encoding/json"
	"net/http"
)

type YouTubeVideoRequest struct {
	ContentID    string `json:"content_id"`
	Likes        int32  `json:"likes"`
	Dislikes     int32  `json:"dislikes"`
	Views        int32  `json:"views"`
	CommentCount int32  `json:"comment_count"`
}

func (a *API) AddYouTubeVideo(writer http.ResponseWriter, request *http.Request) {
	var videoRequest YouTubeVideoRequest
	err := json.NewDecoder(request.Body).Decode(&videoRequest)
	if err != nil {
		writer.WriteHeader(500)
		return
	}

	go func() {
		_ = a.youtubeDislikeRepo.AddDislike(context.Background(), videoRequest.ContentID, videoRequest.Likes, videoRequest.Dislikes, videoRequest.Views, videoRequest.CommentCount)
	}()
	writer.WriteHeader(200)
}
