// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"time"
)

type OpenYoutubeDislikesAggregateDislike struct {
	ID        string    `json:"id"`
	Count     int32     `json:"count"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OpenYoutubeDislikesComment struct {
	VideoID  string  `json:"video_id"`
	Content  string  `json:"content"`
	Negative float32 `json:"negative"`
	Neutral  float32 `json:"neutral"`
	Positive float32 `json:"positive"`
	Compound float32 `json:"compound"`
}

type OpenYoutubeDislikesDislike struct {
	VideoID   string    `json:"video_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type OpenYoutubeDislikesLike struct {
	VideoID   string    `json:"video_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type OpenYoutubeDislikesUser struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type OpenYoutubeDislikesVideo struct {
	ID          string        `json:"id"`
	IDHash      string        `json:"id_hash"`
	Likes       int64         `json:"likes"`
	Dislikes    int64         `json:"dislikes"`
	Views       int64         `json:"views"`
	Comments    sql.NullInt64 `json:"comments"`
	Subscribers int64         `json:"subscribers"`
	PublishedAt int64         `json:"published_at"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DurationSec int32         `json:"duration_sec"`
}
