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

func (d *SingleDislikeRepo) AddDislike(ctx context.Context, contentID string, hashedIP string) error {
	return d.querier.AddDislike(ctx, db.AddDislikeParams{
		ContentID: contentID,
		HashedIp:  hashedIP,
	})
}

func (d SingleDislikeRepo) GetDislikeCount(ctx context.Context, contentID string) (int64, error) {
	return d.querier.GetDislikeCount(ctx, contentID)
}
