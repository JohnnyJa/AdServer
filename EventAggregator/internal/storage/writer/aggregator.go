package writer

import (
	"context"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/mapper"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/model"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/druid-repo"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Aggregator struct {
	buffer     []model.Event
	mutex      sync.Mutex
	flushEvery time.Duration
	maxSize    int
	druid      *druid.Repo
}

func NewAggregator(druid *druid.Repo, flushEvery time.Duration, maxSize int) *Aggregator {
	return &Aggregator{
		buffer:     make([]model.Event, 0),
		mutex:      sync.Mutex{},
		flushEvery: flushEvery,
		maxSize:    maxSize,
		druid:      druid,
	}
}

func (a *Aggregator) Run(eventChan <-chan model.Event) {
	ticker := time.NewTicker(a.flushEvery)

	for {
		select {
		case evt := <-eventChan:
			a.mutex.Lock()
			a.buffer = append(a.buffer, evt)
			if len(a.buffer) >= a.maxSize {
				a.flush()
			}
			a.mutex.Unlock()
		case <-ticker.C:
			a.mutex.Lock()
			a.flush()
			a.mutex.Unlock()
		}
	}
}

func (a *Aggregator) flush() {
	if len(a.buffer) == 0 {
		return
	}
	toWrite := make([]druid.Event, len(a.buffer))

	for i, evt := range a.buffer {
		toWrite[i] = mapper.MapToDruidEvent(evt)
	}

	a.buffer = nil
	go func() {
		err := a.druid.WriteBatch(context.Background(), toWrite)
		if err != nil {
			logrus.Error(err)
			return
		}

	}()
}
