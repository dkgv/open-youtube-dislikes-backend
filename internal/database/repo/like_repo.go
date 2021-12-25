package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type LikeRepo struct {
	querier db.Querier
}

func NewLikeRepo(conn *sql.DB) *LikeRepo {
	return &LikeRepo{querier: db.New(conn)}
}

func (d *LikeRepo) Insert(ctx context.Context, videoID string, userID string) error {
	return d.querier.InsertLike(ctx, db.InsertLikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}

func (d *LikeRepo) Delete(ctx context.Context, videoID string, userID string) error {
	return d.querier.DeleteLike(ctx, db.DeleteLikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}

func (d *LikeRepo) GetLikeCount(ctx context.Context, videoID string) (int64, error) {
	return d.querier.GetLikeCount(ctx, videoID)
}

func (d *LikeRepo) FindByID(ctx context.Context, videoID string, userID string) (db.OpenYoutubeDislikesLike, error) {
	return d.querier.FindLike(ctx, db.FindLikeParams{
		VideoID: videoID,
		UserID:  userID,
	})
}
