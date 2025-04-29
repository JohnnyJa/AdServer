package writer

import (
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/model"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/druid-repo"
	"github.com/sirupsen/logrus"
)

type Pool struct {
	config      *Config
	logger      *logrus.Logger
	druid       *druid.Repo
	ch          <-chan model.Event
	aggregators []*Aggregator
}

func NewPool(config *Config, logger *logrus.Logger, druid *druid.Repo, ch <-chan model.Event) *Pool {
	return &Pool{
		config: config,
		logger: logger,
		druid:  druid,
		ch:     ch,
	}
}

func (p *Pool) Start() error {
	p.aggregators = make([]*Aggregator, p.config.NumOfWriters)
	for i := 0; i < p.config.NumOfWriters; i++ {
		p.aggregators[i] = NewAggregator(p.druid, p.config.FlushTimeout, p.config.MaxSize)
		p.aggregators[i].Run(p.ch)
	}
	return nil
}
