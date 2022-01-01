package swagger

import (
	"flag"
	"log"

	"github.com/dkgv/dislikes/generated/restapi/restapi"
	"github.com/dkgv/dislikes/generated/restapi/restapi/operations"
	"github.com/dkgv/dislikes/internal/logic/dislikes"
	"github.com/dkgv/dislikes/internal/logic/user"
	"github.com/dkgv/dislikes/internal/swagger/handlers"
	"github.com/go-openapi/loads"
)

type API struct {
	DislikeService *dislikes.Service
	UserService    *user.Service
}

func New(dislikeService *dislikes.Service, userService *user.Service) *API {
	return &API{
		DislikeService: dislikeService,
		UserService:    userService,
	}
}

func (a *API) Run() {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOpenYoutubeDislikesBackendAPI(spec)
	handlers.Initialize(a.DislikeService, a.UserService, api)

	server := restapi.NewServer(api)
	server.Port = 5000
	defer server.Shutdown()

	flag.Parse()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
