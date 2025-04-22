package server

import (
	"context"
	"github.com/JohnnyJa/AdServer/EventCollector/container/router"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	redis  *redis.Client
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureRouter(); err != nil {
		return err
	}

	if err := s.configureRedis(); err != nil {
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

	return nil
}

func (s *Server) configureRedis() error {

	opts, err := redis.ParseURL(s.config.RedisConfig.ConnectionString)
	if err != nil {
		return err
	}

	r := redis.NewClient(opts)

	if err := r.Ping(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}

func (s *Server) configureRouter() error {
	r := router.New()
	s.router = r
	return nil
}
