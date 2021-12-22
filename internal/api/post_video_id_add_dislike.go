package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) PostVideoAddDislikeV1(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	ip := GetIP(request)

	go func() {
		_ = a.dataService.AddDislike(context.Background(), id, ip)
	}()
	writer.WriteHeader(http.StatusOK)
}