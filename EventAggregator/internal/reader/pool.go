package reader

import (
	"encoding/json"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/model"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/redis"
	"github.com/sirupsen/logrus"
)

type Pool struct {
	Config *Config
	logger *logrus.Logger
	store  *redis.Redis
	ch     chan model.Event
}

func NewPool(config *Config, store *redis.Redis, logger *logrus.Logger) *Pool {
	ch := make(chan model.Event, config.BufferSize)
	return &Pool{
		Config: config,
		store:  store,
		logger: logger,
		ch:     ch,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.Config.NumOfReaders; i++ {
		go p.startReader()
	}
}

func (p *Pool) Stop() {
	close(p.ch)
}

func (p *Pool) startReader() {
	for {
		rawRequest, err := p.store.Read()
		if err != nil {
			p.logger.Errorf("Failed to read request: %v", err)
		}
		var request model.Event
		err = json.Unmarshal(rawRequest, &request)
		if err != nil {
			p.logger.Errorf("Failed to unmarshal request: %v", err)
		}

		p.logger.Debugf("Read request: %v", request)

		p.ch <- request
	}
}

func (p *Pool) GetChannel() chan model.Event {
	return p.ch
}
