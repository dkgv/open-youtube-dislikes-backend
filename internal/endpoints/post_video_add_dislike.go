package endpoints

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) PostVideoAddDislike(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		_ = a.dataService.AddDislike(context.Background(), id)
	}()
	writer.WriteHeader(http.StatusOK)
}
