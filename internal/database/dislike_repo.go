package database

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type DislikeRepo struct {
	querier db.Querier
}

func NewDislikeRepo(conn *sql.DB) *DislikeRepo {
	return &DislikeRepo{querier: db.New(conn)}
}

func (d *DislikeRepo) AddDislike(ctx context.Context, contentID string, hashedIP string) error {
	return d.querier.AddDislike(ctx, db.AddDislikeParams{
		ContentID: contentID,
		HashedIp:  hashedIP,
	})
}

func (d DislikeRepo) GetDislikeCount(ctx context.Context, contentID string) (int64, error) {
	return d.querier.GetDislikeCount(ctx, contentID)
}
