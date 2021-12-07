// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	AddDislike(ctx context.Context, arg AddDislikeParams) error
	AddYouTubeVideo(ctx context.Context, arg AddYouTubeVideoParams) error
	GetAggregateDislikeCount(ctx context.Context, contentID string) (int32, error)
	GetDislikeCount(ctx context.Context, contentID string) (int64, error)
	SetAggregateDislikeCount(ctx context.Context, arg SetAggregateDislikeCountParams) error
	UpdateAggregateDislikeCount(ctx context.Context, arg UpdateAggregateDislikeCountParams) error
}

var _ Querier = (*Queries)(nil)
