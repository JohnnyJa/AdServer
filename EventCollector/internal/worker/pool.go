package worker

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/store"
	"github.com/sirupsen/logrus"
)

type Pool struct {
	store  *store.Store
	logger *logrus.Logger
	ch     chan interface{}
}

func NewPool(store *store.Store, logger *logrus.Logger) *Pool {
	return &Pool{
		store:  store,
		logger: logger,
		ch:     make(chan interface{}),
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

func (p *Pool) write(data interface{}) {
	err := p.store.Set("request", data)
	if err != nil {
		p.logger.Errorf("Failed to store request: %v", err)
		//TODO: Create recovery logic
	}
}

func (p *Pool) Stop() error {
	close(p.ch)
	return nil
}

func (p *Pool) Write(data interface{}) {
	p.ch <- data
}
