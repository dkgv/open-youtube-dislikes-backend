package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type PostVideoLikeRequest struct {
	Action string `json:"action"`
}

func (a *API) PostVideoLike(writer http.ResponseWriter, request *http.Request) {
	var requestPayload PostVideoLikeRequest
	err := json.NewDecoder(request.Body).Decode(&requestPayload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := GetUserID(request)
	if userID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	videoID, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		if requestPayload.Action == "add" {
			_ = a.dataService.AddLike(context.Background(), videoID, userID)
		} else if requestPayload.Action == "remove" {
			_ = a.dataService.RemoveLike(context.Background(), videoID, userID)
		}
	}()
	writer.WriteHeader(http.StatusOK)
}
