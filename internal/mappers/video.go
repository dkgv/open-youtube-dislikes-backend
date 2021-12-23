package mappers

import (
	db "github.com/dkgv/dislikes/generated/sql"
	"github.com/dkgv/dislikes/internal/types"
)

func DBVideoToVideo(video db.Video) types.Video {
	return types.Video{
		IDHash:      video.IDHash,
		Views:       uint32(video.Views),
		Likes:       uint32(video.Likes),
		Dislikes:    uint32(video.Dislikes),
		Comments:    uint32(video.Comments),
		PublishedAt: uint32(video.PublishedAt),
		Subscribers: uint32(video.Subscribers),
	}
}

func VideoToDBVideo(video types.Video) db.Video {
	return db.Video{
		ID:          video.IDHash,
		IDHash:      video.IDHash,
		Likes:       int64(video.Likes),
		Dislikes:    int64(video.Dislikes),
		Views:       int64(video.Views),
		Comments:    int64(video.Comments),
		Subscribers: int64(video.Subscribers),
	}
}
