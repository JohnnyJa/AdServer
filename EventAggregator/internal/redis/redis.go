package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

type Redis struct {
	redis  *redis.Client
	config *Config
	logger *logrus.Logger
}

func New(config *Config, logger *logrus.Logger) *Redis {
	return &Redis{
		config: config,
		logger: logger,
	}
}

func (s *Redis) Start() error {

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

func (s *Redis) Stop() error {
	err := s.redis.Close()
	if err != nil {
		return err
	}
	s.logger.Info("Redis stopped")
	return nil
}

func (s *Redis) Set(key string, value interface{}) error {
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

func (s *Redis) Ping() error {
	if err := s.redis.Ping(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}

func (s *Redis) Read() ([]byte, error) {
	res, err := s.redis.BLPop(context.Background(), time.Second*0, "request").Result()
	if err != nil {
		return nil, err
	}
	return []byte(res[1]), nil
}
