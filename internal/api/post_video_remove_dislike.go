package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) PostVideoRemoveDislike(writer http.ResponseWriter, request *http.Request) {
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

	go func() {
		_ = a.dataService.RemoveDislike(context.Background(), videoID, userID)
	}()
	writer.WriteHeader(http.StatusOK)
}
