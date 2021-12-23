package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type SingleDislikeRepo struct {
	querier db.Querier
}

func NewSingleDislikeRepo(conn *sql.DB) *SingleDislikeRepo {
	return &SingleDislikeRepo{querier: db.New(conn)}
}

func (d *SingleDislikeRepo) Insert(ctx context.Context, videoID string, userID string) error {
	return d.querier.InsertDislike(ctx, db.InsertDislikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}

func (d *SingleDislikeRepo) Delete(ctx context.Context, videoID string, userID string) error {
	return d.querier.DeleteDislike(ctx, db.DeleteDislikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}

func (d SingleDislikeRepo) GetDislikeCount(ctx context.Context, videoID string) (int64, error) {
	return d.querier.GetDislikeCount(ctx, videoID)
}
