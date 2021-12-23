package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/gorilla/mux"
)

type PostVideoRequest struct {
	types.Video
}

type PostVideoResponse struct {
	HasDisliked bool `json:"has_disliked"`
}

func (a *API) PostVideoV1(writer http.ResponseWriter, request *http.Request) {
	userID := request.Header.Get("X-User-ID")
	if userID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	videoID, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var videoRequest PostVideoRequest
	err := json.NewDecoder(request.Body).Decode(&videoRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		_ = a.dataService.AddVideo(context.Background(), videoID, videoRequest.Video)
	}()

	hasDisliked, err := a.dataService.HasDislikedVideo(context.Background(), userID, videoID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responsePayload, err := json.Marshal(PostVideoResponse{HasDisliked: hasDisliked})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(responsePayload)
}
