package ml

import (
	"context"

	"github.com/dkgv/dislikes/internal/types"
)

type Service struct {
}

func New() (*Service, error) {

	return &Service{}, nil
}

func (s *Service) Predict(ctx context.Context, details types.VideoDetails) (int64, error) {
	return 0, nil
}
