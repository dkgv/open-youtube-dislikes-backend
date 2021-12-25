package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type VideoRepo struct {
	querier db.Querier
}

func NewVideoRepo(conn *sql.DB) *VideoRepo {
	return &VideoRepo{querier: db.New(conn)}
}

func (v *VideoRepo) FindByID(ctx context.Context, id string) (db.OpenYoutubeDislikesVideo, error) {
	return v.querier.FindVideoDetailsByID(ctx, id)
}

func (v *VideoRepo) FindNByHash(ctx context.Context, idHash string, maxCount int32) ([]db.OpenYoutubeDislikesVideo, error) {
	return v.querier.FindNVideosByIDHash(ctx, db.FindNVideosByIDHashParams{
		IDHash: idHash,
		Limit:  maxCount,
	})
}

func (v *VideoRepo) Upsert(ctx context.Context, id string, idHash string, likes, dislikes, views, comments, subscribers, publishedAt uint32) error {
	return v.querier.UpsertVideoDetails(ctx, db.UpsertVideoDetailsParams{
		ID:          id,
		IDHash:      idHash,
		Likes:       int64(likes),
		Dislikes:    int64(dislikes),
		Views:       int64(views),
		Comments:    int64(comments),
		Subscribers: int64(subscribers),
		PublishedAt: int64(publishedAt),
	})
}
