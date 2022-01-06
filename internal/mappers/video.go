package mappers

import (
	"github.com/dkgv/dislikes/generated/restapi/models"
	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/types"
)

func DBVideoToVideo(video db.OpenYoutubeDislikesVideo) types.Video {
	comments := video.Comments.Int64
	return types.Video{
		IDHash:      video.IDHash,
		Views:       video.Views,
		Likes:       video.Likes,
		Dislikes:    video.Dislikes,
		Comments:    comments,
		PublishedAt: video.PublishedAt,
		Subscribers: video.Subscribers,
	}
}

func SwaggerVideoToVideo(video *models.Video) types.Video {
	comments := video.Comments
	return types.Video{
		IDHash:      video.IDHash,
		Views:       video.Views,
		Likes:       video.Likes,
		Dislikes:    video.Dislikes,
		Comments:    comments,
		PublishedAt: video.PublishedAt,
		Subscribers: video.Subscribers,
	}
}
