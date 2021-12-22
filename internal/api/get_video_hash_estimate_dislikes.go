package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dkgv/dislikes/internal/types"
)

type GetVideoHashEstimateDislikesRequest struct {
	types.Video
}

type GetVideoHashEstimateDislikesResponse struct {
	Estimations []types.DislikeEstimation `json:"estimations"`
}

func (a *API) GetVideoHashEstimateDislikesV1(writer http.ResponseWriter, request *http.Request) {
	var requestPayload GetVideoHashEstimateDislikesRequest
	err := json.NewDecoder(request.Body).Decode(&requestPayload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	estimations, err := a.dataService.GetDislikeEstimationsByHash(context.Background(), 1, requestPayload.Video, 5)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responsePayload, err := json.Marshal(GetVideoHashEstimateDislikesResponse{Estimations: estimations})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(responsePayload)
}
