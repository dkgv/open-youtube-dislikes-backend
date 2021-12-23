package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dkgv/dislikes/internal/logic/data"
	"github.com/gorilla/mux"
)

type API struct {
	dataService *data.Service
}

func New(dataService *data.Service) *API {
	return &API{
		dataService: dataService,
	}
}

func (a *API) Run() {
	router := mux.NewRouter()

	router.HandleFunc(apiURL(1, "video/{id}"), a.PostVideoV1).Methods("POST")
	router.HandleFunc(apiURL(1, "video/{id}/add_dislike"), a.PostVideoAddDislike).Methods("POST")
	router.HandleFunc(apiURL(1, "video/{id}/remove_dislike"), a.PostVideoAddDislike).Methods("POST")
	router.HandleFunc(apiURL(1, "video/{id}/dislikes"), a.GetVideoEstimateDislikesV1).Methods("GET")
	router.HandleFunc(apiURL(1, "video/{id_hash}/dislikes"), a.GetVideoHashEstimateDislikesV1).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func apiURL(version int, endpoint string) string {
	return fmt.Sprintf("/api/v%d%s", version, endpoint)
}

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) == 0 {
		ip = r.Header.Get("CF-Connecting-IP")
	}
	if len(ip) == 0 {
		ip = r.Header.Get("X-Real-IP")
	}
	if len(ip) == 0 {
		ip = r.RemoteAddr
	}
	return ip
}

func GetUserID(r *http.Request) string {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		return ""
	}

	if len(userID) != 36 {
		return ""
	}

	return userID
}
