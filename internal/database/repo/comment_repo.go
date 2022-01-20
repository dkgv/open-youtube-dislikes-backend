package repo

import (
	"context"
	"database/sql"

	db "github.com/dkgv/dislikes/generated/sql"
)

type CommentRepo struct {
	querier db.Querier
}

func NewCommentRepo(conn *sql.DB) *CommentRepo {
	return &CommentRepo{querier: db.New(conn)}
}

func (c *CommentRepo) Insert(ctx context.Context, videoID string, content string, negative float32, neutral float32, positive float32, compound float32) error {
	return c.querier.InsertComment(ctx, db.InsertCommentParams{
		VideoID:  videoID,
		Content:  content,
		Negative: negative,
		Neutral:  neutral,
		Positive: positive,
		Compound: compound,
	})
}
