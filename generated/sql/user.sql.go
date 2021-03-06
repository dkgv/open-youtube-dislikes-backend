// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package db

import (
	"context"
)

const findUserByID = `-- name: FindUserByID :one
SELECT id, created_at FROM open_youtube_dislikes."user" WHERE id = $1
`

func (q *Queries) FindUserByID(ctx context.Context, id string) (OpenYoutubeDislikesUser, error) {
	row := q.queryRow(ctx, q.findUserByIDStmt, findUserByID, id)
	var i OpenYoutubeDislikesUser
	err := row.Scan(&i.ID, &i.CreatedAt)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO open_youtube_dislikes."user" (id) VALUES ($1) ON CONFLICT DO NOTHING
`

func (q *Queries) InsertUser(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.insertUserStmt, insertUser, id)
	return err
}
