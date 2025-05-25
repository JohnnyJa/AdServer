package server

import (
	"errors"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/kafka"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/router"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router router.Router
	kafka  kafka.Kafka
	client grpcClients.IncrementViewsClient
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

	if err := s.configureKafka(); err != nil {
		return err
	}

	if err := s.configureClient(); err != nil {
		return err
	}

	if err := s.configureRouter(); err != nil {
		return err
	}

	s.logger.Info("Starting API Server")

	if err := s.router.Start(); err != nil {
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

func (s *Server) configureClient() error {

	client := grpcClients.NewIncrementViewsClient(s.config.ClientConfig)

	err := client.Start()
	if err != nil {
		return err
	}

	s.client = client
	return nil
}

func (s *Server) configureRouter() error {

	if s.logger == nil {
		return errors.New("no logger configured")
	}
	if s.kafka == nil {
		return errors.New("no kafka configured")
	}

	r := router.New(s.config.AppConfig.Port, s.logger, s.kafka, s.client)
	err := r.ConfigureRoute()
	if err != nil {
		return err
	}

	err = r.Start()
	if err != nil {
		return err
	}

	s.router = r

	s.logger.Info("Router configured")
	return nil
}

func (s *Server) configureKafka() error {
	if s.logger == nil {
		return errors.New("no logger configured")
	}

	k := kafka.New(s.config.KafkaConfig, s.logger)
	err := k.Start()
	if err != nil {
		return err
	}

	s.kafka = k

	s.logger.Info("Kafka configured")
	return nil
}
