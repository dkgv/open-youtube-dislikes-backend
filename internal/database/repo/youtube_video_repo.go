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

func (d *YouTubeVideoRepo) Upsert(ctx context.Context, id string, likes, dislikes, views, comments int64) error {
	return d.querier.UpsertYouTubeVideo(ctx, db.UpsertYouTubeVideoParams{
		ID:       id,
		Likes:    likes,
		Dislikes: dislikes,
		Views:    views,
		Comments: comments,
	})
}
