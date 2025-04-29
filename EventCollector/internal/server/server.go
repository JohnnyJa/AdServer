package server

import (
	"errors"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/redis"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/router"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/worker"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	config  *Config
	logger  *logrus.Logger
	router  *router.Router
	redis   *redis.Redis
	workers *worker.Pool
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	if err := s.configurePool(); err != nil {
		return err
	}

	if err := s.configureRouter(); err != nil {
		return err
	}

	s.logger.Info("Starting API Server")

	if err := s.router.Run(s.config.AppConfig.Port); err != nil {
		return err
	}
	return nil
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.AppConfig.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	s.logger.SetOutput(os.Stdout)
	s.logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	s.logger.Info("Logger configured")

	return nil
}

func (s *Server) configureStore() error {
	st := redis.New(s.config.RedisConfig, s.logger)

	err := st.Start()
	if err != nil {
		return err
	}

	s.redis = st

	s.logger.Info("Redis configured")
	return nil
}

func (s *Server) configurePool() error {
	p := worker.NewPool(s.redis, s.logger)
	p.Start(10)
	s.workers = p
	return nil
}

func (s *Server) configureRouter() error {

	if s.workers == nil {
		return errors.New("no workers configured")
	}

	if s.logger == nil {
		return errors.New("no logger configured")
	}

	r := router.New(s.workers, s.logger)
	r.Start()

	s.router = r

	s.logger.Info("Router configured")
	return nil
}
