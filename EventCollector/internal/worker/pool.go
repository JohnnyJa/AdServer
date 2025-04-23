package worker

import (
	"encoding/json"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/model"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/store"
	"github.com/sirupsen/logrus"
)

type Pool struct {
	store  *store.Store
	logger *logrus.Logger
	ch     chan model.Event
}

func NewPool(store *store.Store, logger *logrus.Logger) *Pool {
	return &Pool{
		store:  store,
		logger: logger,
		ch:     make(chan model.Event),
	}
}

func (p *Pool) Start(n int) {
	go func() {
		for i := 0; i < n; i++ {
			for rq := range p.ch {
				p.write(rq)
			}
		}
	}()
}

func (p *Pool) write(data model.Event) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		p.logger.Fatal("Cannot marshal event to json: %v", err)
		return
	}

	err = p.store.Set("request", jsonData)
	if err != nil {
		p.logger.Errorf("Failed to store request: %v", err)
		//TODO: Create recovery logic
	}
}

func (p *Pool) Stop() error {
	close(p.ch)
	return nil
}

func (p *Pool) Write(data model.Event) {
	p.ch <- data
}
