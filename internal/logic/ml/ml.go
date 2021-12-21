package ml

import (
	"context"

	xgb "github.com/Elvenson/xgboost-go"
	"github.com/Elvenson/xgboost-go/activation"
	"github.com/Elvenson/xgboost-go/inference"
	"github.com/Elvenson/xgboost-go/mat"
	"github.com/dkgv/dislikes/internal/types"
)

type Service struct {
	ensemble *inference.Ensemble
}

func New() (*Service, error) {
	ensemble, err := xgb.LoadXGBoostFromJSON(
		"statics/xgboost.json",
		"",
		1,
		0,
		&activation.Raw{},
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		ensemble: ensemble,
	}, nil
}

func (s *Service) Predict(ctx context.Context, details types.VideoDetails) (int64, error) {
	input := mat.SparseMatrix{
		Vectors: []mat.SparseVector{
			{
				0: float64(details.Views),
				1: float64(details.Likes),
				2: details.LikePerView(),
			},
		},
	}

	output, err := s.ensemble.PredictRegression(input, 0)
	if err != nil || len(output.Vectors) == 0 {
		return 0, err
	}

	// Output should be a single number
	vector := output.Vectors[0]
	floats := []float64(*vector)
	if len(floats) == 0 {
		return 0, ErrNoPrediction
	}

	return int64(floats[0]), nil
}
