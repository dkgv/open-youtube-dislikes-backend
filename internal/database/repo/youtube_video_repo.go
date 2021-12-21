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

func (y *YouTubeVideoRepo) FindByID(ctx context.Context, id string) (db.YoutubeVideo, error) {
	return y.querier.FindYouTubeVideoByID(ctx, id)
}

func (y *YouTubeVideoRepo) Upsert(ctx context.Context, id string, likes, dislikes, views, comments int64) error {
	return y.querier.UpsertYouTubeVideo(ctx, db.UpsertYouTubeVideoParams{
		ID:       id,
		Likes:    likes,
		Dislikes: dislikes,
		Views:    views,
		Comments: comments,
	})
}
