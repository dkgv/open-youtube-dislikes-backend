package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type YouTubeVideoRepo struct {
	querier db.Querier
}

func NewYouTubeVideoRepo(conn *sql.DB) *YouTubeVideoRepo {
	return &YouTubeVideoRepo{querier: db.New(conn)}
}

func (d *YouTubeVideoRepo) AddDislike(ctx context.Context, contentID string, likes, dislikes, views, commentCount int32) error {
	return d.querier.AddYouTubeVideo(ctx, db.AddYouTubeVideoParams{
		ContentID:    contentID,
		Likes:        likes,
		Dislikes:     dislikes,
		Views:        views,
		CommentCount: commentCount,
	})
}
