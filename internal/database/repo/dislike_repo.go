package repo

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

func (d *DislikeRepo) Insert(ctx context.Context, videoID string, userID string) error {
	return d.querier.InsertDislike(ctx, db.InsertDislikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}

func (d *DislikeRepo) Delete(ctx context.Context, videoID string, userID string) error {
	return d.querier.DeleteDislike(ctx, db.DeleteDislikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}

func (d *DislikeRepo) GetDislikeCount(ctx context.Context, videoID string) (int64, error) {
	return d.querier.GetDislikeCount(ctx, videoID)
}

func (d *DislikeRepo) FindByID(ctx context.Context, videoID string, userID string) (db.OpenYoutubeDislikesDislike, error) {
	return d.querier.FindDislike(ctx, db.FindDislikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}
