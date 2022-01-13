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

func (s *Service) Predict(ctx context.Context, apiVersion int, video types.Video) (int64, error) {
	switch apiVersion {
	case 1:
		return s.predictV1(video)

	default:
		return 0, ErrUnsupportedAPIVersion
	}
}

func (s *Service) predictV1(video types.Video) (int64, error) {
	input := mat.SparseMatrix{
		Vectors: []mat.SparseVector{
			{
				0:  float64(video.Views),
				1:  float64(video.Likes),
				2:  float64(video.Comments),
				3:  float64(video.Subscribers),
				4:  float64(video.DaysSincePublish()),
				5:  float64(video.DurationSec),
				6:  video.ViewsPerLike(),
				7:  video.LikesPerComment(),
				8:  video.ViewsPerComment(),
				9:  video.DaysPerLike(),
				10: video.DaysPerComment(),
			},
		},
	}

	return tryPredict(s.modelV1, input)
}

func tryPredict(model *inference.Ensemble, input mat.SparseMatrix) (int64, error) {
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

	return int64(floats[0]), nil
}
