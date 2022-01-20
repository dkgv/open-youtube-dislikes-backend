package ml

import (
	"context"

	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func (s *Service) Sentiment(ctx context.Context, text string) (sentitext.Sentiment, error) {
	sentiText := sentitext.Parse(text, lexicon.DefaultLexicon)
	score := sentitext.PolarityScore(sentiText)
	return score, nil
}
