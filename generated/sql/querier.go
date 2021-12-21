// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	FindAggregateDislikeByID(ctx context.Context, id string) (int32, error)
	FindYouTubeVideoByID(ctx context.Context, id string) (YoutubeVideo, error)
	GetDislikeCount(ctx context.Context, id string) (int64, error)
	InsertAggregateDislike(ctx context.Context, arg InsertAggregateDislikeParams) error
	InsertDislike(ctx context.Context, arg InsertDislikeParams) error
	UpdateAggregateDislike(ctx context.Context, arg UpdateAggregateDislikeParams) error
	UpsertYouTubeVideo(ctx context.Context, arg UpsertYouTubeVideoParams) error
}

var _ Querier = (*Queries)(nil)
