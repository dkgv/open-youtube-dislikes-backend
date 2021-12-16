package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dkgv/dislikes/internal/types"
	"github.com/gorilla/mux"
)

type PostVideoRequest struct {
	types.VideoDetails
}

func (a *API) PostVideoAdd(writer http.ResponseWriter, request *http.Request) {
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
		_ = a.dataService.AddYouTubeVideo(context.Background(), id, videoRequest.VideoDetails)
	}()
	writer.WriteHeader(http.StatusOK)
}
