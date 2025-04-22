package store

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Store struct {
	redis  *redis.Client
	config *Config
	logger *logrus.Logger
}

func New(config *Config, logger *logrus.Logger) *Store {
	return &Store{
		config: config,
		logger: logger,
	}
}

func (s *Store) Start() error {

	opts, err := redis.ParseURL(s.config.ConnectionString)
	if err != nil {
		return err
	}

	r := redis.NewClient(opts)

	if err := r.Ping(context.Background()).Err(); err != nil {
		return err
	}

	s.redis = r

	s.logger.Info("Redis configured")
	return nil
}

func (s *Store) Stop() error {
	err := s.redis.Close()
	if err != nil {
		return err
	}
	s.logger.Info("Redis stopped")
	return nil
}

func (s *Store) Set(key string, value interface{}) error {
	if err := s.redis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: key,
		Values: map[string]interface{}{
			"value": value,
		},
	}).Err(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Ping() error {
	if err := s.redis.Ping(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}
