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

func (d *AggregateDislikeRepo) GetAggregateDislikeCount(ctx context.Context, contentID string) (int32, error) {
	return d.querier.GetAggregateDislikeCount(ctx, contentID)
}

func (d *AggregateDislikeRepo) SetAggregateDislikeCount(ctx context.Context, contentID string, count int32) error {
	return d.querier.SetAggregateDislikeCount(ctx, db.SetAggregateDislikeCountParams{
		ContentID: contentID,
		Count:     count,
	})
}
