package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dkgv/dislikes/internal/types"
)

type GetVideoEstimateDislikesRequest struct {
	types.VideoDetails
}

func (a *API) GetVideoEstimateDislikes(writer http.ResponseWriter, request *http.Request) {
	var statsRequest GetVideoEstimateDislikesRequest
	err := json.NewDecoder(request.Body).Decode(&statsRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	stats, err := a.dataService.GetDislikes(context.Background(), statsRequest.VideoDetails)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	obj, err := json.Marshal(stats)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(obj)
}
