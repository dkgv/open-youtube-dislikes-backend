package handlers

import (
	"context"
	"log"

	"github.com/dkgv/dislikes/generated/restapi/models"
	"github.com/dkgv/dislikes/generated/restapi/restapi/operations"
	"github.com/dkgv/dislikes/internal/logic/data"
	"github.com/dkgv/dislikes/internal/logic/user"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/go-openapi/runtime/middleware"
)

func Initialize(dataService *data.Service, userService *user.Service, swagger *operations.OpenYoutubeDislikesBackendAPI) {
	swagger.PostVideoIDHandler = operations.PostVideoIDHandlerFunc(func(params operations.PostVideoIDParams) middleware.Responder {
		ctx := context.Background()

		if params.Video == nil {
			log.Println("Video is nil")
			return operations.NewPostVideoIDBadRequest()
		}


		err := dataService.AddVideo(context.Background(), params.ID, mappers.SwaggerVideoToVideo(params.Video))
		if err != nil {
			log.Println("Error while adding video: ", err)
			return operations.NewPostVideoIDBadRequest()
		}

		hasLiked, err := userService.HasLikedVideo(ctx, params.ID, params.XUserID)
		if err != nil {
			log.Println("Erro likedr while adding video: ", err)
			return operations.NewPostVideoIDBadRequest()
		}

		hasDisliked, err := userService.HasDislikedVideo(ctx, params.ID, params.XUserID)
		if err != nil {
			log.Println("Erro dislikedr while adding video: ", err)
			return operations.NewPostVideoIDBadRequest()
		}

		dislikes, formattedDislikes, err := dataService.GetDislikes(ctx, 1, mappers.SwaggerVideoToVideo(params.Video))
		if err != nil {
			log.Println("Erro dislikes get while adding video: ", err)
			return operations.NewPostVideoIDBadRequest()
		}

		return operations.NewPostVideoIDOK().WithPayload(&models.VideoResponse{
			HasDisliked:       hasDisliked,
			HasLiked:          hasLiked,
			Dislikes:          dislikes,
			FormattedDislikes: formattedDislikes,
		})
	})

	swagger.PostVideoIDLikeHandler = operations.PostVideoIDLikeHandlerFunc(func(params operations.PostVideoIDLikeParams) middleware.Responder {
		go func() {
			if params.Action == "add" {
				_ = userService.AddLike(context.Background(), params.ID, params.XUserID)
			} else {
				_ = userService.RemoveLike(context.Background(), params.ID, params.XUserID)
			}
		}()

		return operations.NewPostVideoIDLikeOK()
	})

	swagger.PostVideoIDDislikeHandler = operations.PostVideoIDDislikeHandlerFunc(func(params operations.PostVideoIDDislikeParams) middleware.Responder {
		go func() {
			if params.Action == "add" {
				_ = userService.AddDislike(context.Background(), params.ID, params.XUserID)
			} else {
				_ = userService.RemoveDislike(context.Background(), params.ID, params.XUserID)
			}
		}()

		return operations.NewPostVideoIDDislikeOK()
	})
}
