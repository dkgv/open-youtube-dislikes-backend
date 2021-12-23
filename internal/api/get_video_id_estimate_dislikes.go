package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dkgv/dislikes/internal/types"
)

type GetVideoEstimateDislikesRequest struct {
	types.Video
}

type GetVideoEstimateDislikesResponse struct {
	Dislikes string `json:"dislikes"`
}

func (a *API) GetVideoEstimateDislikesV1(writer http.ResponseWriter, request *http.Request) {
	a.getVideoEstimateDislikes(writer, request, 1)
}

func (a *API) getVideoEstimateDislikes(writer http.ResponseWriter, request *http.Request, apiVersion int) {
	var requestPayload GetVideoEstimateDislikesRequest
	err := json.NewDecoder(request.Body).Decode(&requestPayload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	dislikes, err := a.dataService.GetDislikes(context.Background(), apiVersion, requestPayload.Video)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responsePayload, err := json.Marshal(GetVideoEstimateDislikesResponse{Dislikes: dislikes})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(responsePayload)
}
