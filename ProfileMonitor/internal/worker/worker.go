package worker

import (
	"context"
	"fmt"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/kafka"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/mapper"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/repository"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/service"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
	"time"
)

type Worker interface {
	service.Service
}

type worker struct {
	config *Config
	repo   repository.Repository
	kafka  kafka.Kafka
	logger *logrus.Logger

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewWorker(config *Config, logger *logrus.Logger, repo repository.Repository, kafka kafka.Kafka) Worker {
	return &worker{
		config: config,
		repo:   repo,
		kafka:  kafka,
		logger: logger,
	}
}

func (w *worker) Start() error {
	w.ctx, w.cancel = context.WithCancel(context.Background())
	w.wg.Add(1)

	go func() {
		defer w.wg.Done()
		ticker := time.NewTicker(w.config.Delay)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := w.runCycle(w.ctx); err != nil {
					log.Printf("worker runCycle error: %v", err)
				}
			case <-w.ctx.Done():
				log.Println("worker received stop signal")
				return
			}
		}
	}()

	return nil
}

func (w *worker) Stop() error {
	w.cancel()
	w.wg.Wait()
	return nil
}

func (w *worker) runCycle(ctx context.Context) error {
	rows, err := w.repo.ReadProfiles(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active profiles: %w", err)
	}

	grouped, err := mapper.ToProfiles(rows)
	if err != nil {
		return fmt.Errorf("failed to convert to profiles: %w", err)
	}
	messages := make([]kafka.ProfileForKafka, len(grouped))

	i := 0
	for _, profile := range grouped {
		messages[i] = mapper.ToKafkaProfile(*profile)
		i++
	}

	return w.kafka.Write(messages)
}
