package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type AggregateDislikeRepo struct {
	querier db.Querier
}

func NewAggregateDislikeRepo(conn *sql.DB) *AggregateDislikeRepo {
	return &AggregateDislikeRepo{querier: db.New(conn)}
}

func (d *AggregateDislikeRepo) FindByID(ctx context.Context, videoID string) (int32, error) {
	return d.querier.FindAggregateDislikeByID(ctx, videoID)
}

func (d *AggregateDislikeRepo) UpdateByID(ctx context.Context, videoID string, count int32) error {
	return d.querier.UpdateAggregateDislike(ctx, db.UpdateAggregateDislikeParams{
		ID:    videoID,
		Count: count,
	})
}
