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
	modelV1 *inference.Ensemble
}

func New() (*Service, error) {
	modelV1, err := xgb.LoadXGBoostFromJSON(
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
		modelV1: modelV1,
	}, nil
}

func (s *Service) Predict(ctx context.Context, apiVersion int, video types.Video) (uint32, error) {
	switch apiVersion {
	case 1:
		return s.predictV1(video)

	default:
		return 0, ErrUnsupportedAPIVersion
	}
}

func (s *Service) predictV1(video types.Video) (uint32, error) {
	input := mat.SparseMatrix{
		Vectors: []mat.SparseVector{
			{
				0: float64(video.Views),
				1: float64(video.Likes),
				2: video.LikesPerView(),
			},
		},
	}

	return tryPredict(s.modelV1, input)
}

func tryPredict(model *inference.Ensemble, input mat.SparseMatrix) (uint32, error) {
	output, err := model.PredictRegression(input, 0)
	if err != nil || len(output.Vectors) == 0 {
		return 0, err
	}

	// Output should be a single number
	vector := output.Vectors[0]
	floats := []float64(*vector)
	if len(floats) == 0 {
		return 0, ErrNoPrediction
	}

	return uint32(floats[0]), nil
}
