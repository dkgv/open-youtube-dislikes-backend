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
	writer.WriteHeader(http.StatusOK)
}
