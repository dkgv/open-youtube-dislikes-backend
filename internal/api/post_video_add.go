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

func (a *API) PostVideoAddV1(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, ok := vars["id"]
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
		_ = a.dataService.AddVideo(context.Background(), id, videoRequest.Video)
	}()
	writer.WriteHeader(http.StatusOK)
}
