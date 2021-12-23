package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/gorilla/mux"
)

type PostVideoAddDislikeRequest struct {
	types.Video
}

func (a *API) PostVideoAddDislike(writer http.ResponseWriter, request *http.Request) {
	userID := GetUserID(request)
	if userID == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	videoID, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestPayload GetVideoEstimateDislikesRequest
	err := json.NewDecoder(request.Body).Decode(&requestPayload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		_ = a.dataService.AddDislike(context.Background(), videoID, userID)
	}()
	writer.WriteHeader(http.StatusOK)
}
