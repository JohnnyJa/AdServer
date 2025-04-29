package server

import (
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/reader"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/redis"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/druid-repo"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/writer"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	config  *Config
	logger  *logrus.Logger
	redis   *redis.Redis
	readers *reader.Pool
	writers *writer.Pool
	druid   *druid.Repo
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

	if err := s.configureRedis(); err != nil {
		return err
	}

	if err := s.configureReaders(); err != nil {
		return err
	}

	if err := s.configureDruid(); err != nil {
		return err
	}

	if err := s.configureWriter(); err != nil {
		return err
	}
	s.logger.Info("Starting API Server")

	select {}
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

func (s *Server) configureRedis() error {
	st := redis.New(s.config.RedisConfig, s.logger)

	err := st.Start()
	if err != nil {
		return err
	}

	s.redis = st

	s.logger.Info("Store configured")
	return nil
}

func (s *Server) configureReaders() error {
	p := reader.NewPool(s.config.ReaderConfig, s.redis, s.logger)
	p.Start()
	s.readers = p
	return nil
}

func (s *Server) configureDruid() error {
	st := druid.NewRepo(s.config.DruidConfig)
	s.druid = st

	s.logger.Info("Clickhouse configured")
	return nil
}

func (s *Server) configureWriter() error {
	p := writer.NewPool(s.config.WriterConfig, s.logger, s.druid, s.readers.GetChannel())
	err := p.Start()
	if err != nil {
		return err
	}

	s.writers = p
	s.logger.Info("Writers configured")
	return nil

}
