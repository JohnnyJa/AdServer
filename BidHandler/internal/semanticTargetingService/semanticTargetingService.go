package semanticTargetingService

import (
	"context"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/service"
	"github.com/fluhus/gostuff/nlp/wordnet"
	"github.com/sirupsen/logrus"
)

type SemanticTargetingService interface {
	service.Service
	IsSimilar(wordA string, wordB string) bool
}

type semanticTargetingService struct {
	wn     *wordnet.WordNet
	config *app.SemanticTargetingConfig
	logger *logrus.Logger
}

func New(config *app.Config, logger *logrus.Logger) SemanticTargetingService {
	return &semanticTargetingService{
		config: config.SemanticTargetingConfig,
		logger: logger,
	}
}

func (s *semanticTargetingService) Start(ctx context.Context) error {
	parse, err := wordnet.Parse(s.config.Path)
	if err != nil {
		return err
	}

	s.wn = parse
	return nil
}

func (s *semanticTargetingService) Stop(ctx context.Context) error {
	return nil
}

func (s *semanticTargetingService) IsSimilar(wordA string, wordB string) bool {
	wnWordA := s.wn.Search(wordA)["n"][0]
	wnWordB := s.wn.Search(wordB)["n"][0]

	return s.wn.PathSimilarity(wnWordA, wnWordB, false) >= s.config.MinSimilarity
}
