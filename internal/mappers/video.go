package mappers

import (
	"github.com/dkgv/dislikes/generated/restapi/models"
	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/types"
)

func DBVideoToVideo(video db.OpenYoutubeDislikesVideo) types.Video {
	comments := uint32(video.Comments.Int64)
	return types.Video{
		IDHash:      video.IDHash,
		Views:       uint32(video.Views),
		Likes:       uint32(video.Likes),
		Dislikes:    uint32(video.Dislikes),
		Comments:    &comments,
		PublishedAt: video.PublishedAt,
		Subscribers: uint32(video.Subscribers),
	}
}

func SwaggerVideoToVideo(video *models.Video) types.Video {
	comments := uint32(video.Comments)
	return types.Video{
		IDHash:      video.IDHash,
		Views:       uint32(video.Views),
		Likes:       uint32(video.Likes),
		Dislikes:    uint32(video.Dislikes),
		Comments:    &comments,
		PublishedAt: video.PublishedAt,
		Subscribers: uint32(video.Subscribers),
	}
}
