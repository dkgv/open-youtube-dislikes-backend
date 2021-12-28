// Code generated by sqlc. DO NOT EDIT.
// source: video.sql

package db

import (
	"context"
	"database/sql"
)

const findNVideosByIDHash = `-- name: FindNVideosByIDHash :many
SELECT id, id_hash, likes, dislikes, views, comments, subscribers, published_at, created_at, updated_at FROM open_youtube_dislikes.video WHERE id_hash LIKE $1 LIMIT $2
`

type FindNVideosByIDHashParams struct {
	IDHash string `json:"id_hash"`
	Limit  int32  `json:"limit"`
}

func (q *Queries) FindNVideosByIDHash(ctx context.Context, arg FindNVideosByIDHashParams) ([]OpenYoutubeDislikesVideo, error) {
	rows, err := q.query(ctx, q.findNVideosByIDHashStmt, findNVideosByIDHash, arg.IDHash, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OpenYoutubeDislikesVideo{}
	for rows.Next() {
		var i OpenYoutubeDislikesVideo
		if err := rows.Scan(
			&i.ID,
			&i.IDHash,
			&i.Likes,
			&i.Dislikes,
			&i.Views,
			&i.Comments,
			&i.Subscribers,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findVideoDetailsByID = `-- name: FindVideoDetailsByID :one
SELECT id, id_hash, likes, dislikes, views, comments, subscribers, published_at, created_at, updated_at FROM open_youtube_dislikes.video WHERE id = $1
`

func (q *Queries) FindVideoDetailsByID(ctx context.Context, id string) (OpenYoutubeDislikesVideo, error) {
	row := q.queryRow(ctx, q.findVideoDetailsByIDStmt, findVideoDetailsByID, id)
	var i OpenYoutubeDislikesVideo
	err := row.Scan(
		&i.ID,
		&i.IDHash,
		&i.Likes,
		&i.Dislikes,
		&i.Views,
		&i.Comments,
		&i.Subscribers,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertVideoDetails = `-- name: UpsertVideoDetails :exec
INSERT INTO open_youtube_dislikes.video
    (id, id_hash, likes, dislikes, views, comments, subscribers, published_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    ON CONFLICT (id) DO
        UPDATE SET
            likes = excluded.likes,
            dislikes = excluded.dislikes,
            views = excluded.views,
            comments = excluded.comments,
            subscribers = excluded.subscribers
        WHERE video.likes <= excluded.likes
            AND video.dislikes <= excluded.dislikes
            AND video.views < excluded.views
            AND video.comments <= excluded.comments
            AND video.subscribers <= excluded.subscribers
            AND video.published_at = excluded.published_at
`

type UpsertVideoDetailsParams struct {
	ID          string        `json:"id"`
	IDHash      string        `json:"id_hash"`
	Likes       int64         `json:"likes"`
	Dislikes    int64         `json:"dislikes"`
	Views       int64         `json:"views"`
	Comments    sql.NullInt64 `json:"comments"`
	Subscribers int64         `json:"subscribers"`
	PublishedAt int64         `json:"published_at"`
}

func (q *Queries) UpsertVideoDetails(ctx context.Context, arg UpsertVideoDetailsParams) error {
	_, err := q.exec(ctx, q.upsertVideoDetailsStmt, upsertVideoDetails,
		arg.ID,
		arg.IDHash,
		arg.Likes,
		arg.Dislikes,
		arg.Views,
		arg.Comments,
		arg.Subscribers,
		arg.PublishedAt,
	)
	return err
}
