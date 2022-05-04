package ml

import (
	"context"
	"fmt"

	xgb "github.com/Elvenson/xgboost-go"
	"github.com/Elvenson/xgboost-go/activation"
	"github.com/Elvenson/xgboost-go/inference"
	"github.com/Elvenson/xgboost-go/mat"
	"github.com/dkgv/dislikes/internal/types"
)

type ModelType int

const (
	ModelTypeSimple    ModelType = iota
	ModelTypeSentiment ModelType = 1
)

type Service struct {
	simpleModel    *inference.Ensemble
	sentimentModel *inference.Ensemble
}

func New() (*Service, error) {
	simpleModel, err := loadModel("xgboost-simple")
	if err != nil {
		return nil, err
	}

	sentimentModel, err := loadModel("xgboost-sentiment")
	if err != nil {
		return nil, err
	}

	return &Service{
		simpleModel:    simpleModel,
		sentimentModel: sentimentModel,
	}, nil
}

func loadModel(name string) (*inference.Ensemble, error) {
	modelV1, err := xgb.LoadXGBoostFromJSON(
		fmt.Sprintf("statics/%s.json", name),
		"",
		1,
		0,
		&activation.Raw{},
	)
	return modelV1, err
}

func (s *Service) Predict(ctx context.Context, apiVersion ModelType, video types.Video) (int64, error) {
	switch apiVersion {
	case ModelTypeSimple:
		return s.predictSimple(video)

	case ModelTypeSentiment:
		return s.predictSentiment(video)

	default:
		return 0, ErrUnsupportedAPIVersion
	}
}

func (s *Service) predictSimple(video types.Video) (int64, error) {
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

	return tryPredict(s.simpleModel, input)
}

func (s *Service) predictSentiment(video types.Video) (int64, error) {
	input := mat.SparseMatrix{
		Vectors: []mat.SparseVector{
			{
				0:  float64(video.Likes),
				1:  float64(video.Dislikes),
				2:  float64(video.Views),
				3:  float64(video.Comments),
				4:  float64(video.Subscribers),
				5:  float64(video.DaysSincePublish()),
				6:  float64(video.DurationSec),
				7:  video.Negative,
				8:  video.Neutral,
				9:  video.Positive,
				10: video.Compound,
				11: video.ViewsPerLike(),
				12: video.SubscribersPerLike(),
				13: video.LikesPerComment(),
				14: video.ViewsPerComment(),
				15: video.DaysPerLike(),
				16: video.DaysPerComment(),
			},
		},
	}

	return tryPredict(s.sentimentModel, input)
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

	dislikes := int64(floats[0])
	if dislikes < 0 {
		dislikes = 0
	}

	return dislikes, nil
}
