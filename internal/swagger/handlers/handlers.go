package handlers

import (
	"context"
	"log"

	"github.com/dkgv/dislikes/generated/restapi/models"
	"github.com/dkgv/dislikes/generated/restapi/restapi/operations"
	"github.com/dkgv/dislikes/internal/logic/dislikes"
	"github.com/dkgv/dislikes/internal/logic/user"
	"github.com/dkgv/dislikes/internal/logic/video"
	"github.com/dkgv/dislikes/internal/mappers"
	"github.com/go-openapi/runtime/middleware"
)

func Initialize(dislikeService *dislikes.Service, userService *user.Service, videoService *video.Service, swagger *operations.OpenYoutubeDislikesBackendAPI) {
	swagger.PostVideoIDHandler = operations.PostVideoIDHandlerFunc(func(params operations.PostVideoIDParams) middleware.Responder {
		ctx := context.Background()

		if params.Video == nil {
			return operations.NewPostVideoIDBadRequest()
		}

		err := videoService.ProcessVideo(context.Background(), params.ID, mappers.SwaggerVideoToVideo(params.Video))
		if err != nil {
			log.Printf("Error while adding video: %v", err)
			return operations.NewPostVideoIDBadRequest()
		}

		hasLiked, err := userService.HasLikedVideo(ctx, params.ID, params.XUserID)
		if err != nil {
			return operations.NewPostVideoIDBadRequest()
		}

		hasDisliked, err := userService.HasDislikedVideo(ctx, params.ID, params.XUserID)
		if err != nil {
			return operations.NewPostVideoIDBadRequest()
		}

		dislikes, formattedDislikes, err := dislikeService.GetDislikes(ctx, params.ID)
		if err != nil {
			log.Printf("Error while getting dislikes: %v", err)
			return operations.NewPostVideoIDBadRequest()
		}

		return operations.NewPostVideoIDOK().WithPayload(&models.VideoResponse{
			HasDisliked:       &hasDisliked,
			HasLiked:          &hasLiked,
			Dislikes:          &dislikes,
			FormattedDislikes: &formattedDislikes,
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
