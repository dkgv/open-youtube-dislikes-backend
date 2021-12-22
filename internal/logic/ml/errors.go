package ml

import "errors"

var (
	ErrNoPrediction          = errors.New("NO_PREDICTION")
	ErrUnsupportedAPIVersion = errors.New("UNSUPPORTED_API_VERSION")
)
