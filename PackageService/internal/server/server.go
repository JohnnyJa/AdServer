package server

import (
	"errors"
	"github.com/JohnnyJa/AdServer/PackageService/internal/gRPC"
	"github.com/JohnnyJa/AdServer/PackageService/internal/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

type Server struct {
	config     *Config
	logger     *logrus.Logger
	repo       repository.Repository
	grpcServer *grpc.Server
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

	if err := s.configureRepository(); err != nil {
		return err
	}

	s.logger.Info("Starting API Server")

	if err := s.startGRPCServer(); err != nil {
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

func (s *Server) configureRepository() error {
	if s.logger == nil {
		return errors.New("no logger configured")
	}

	r := repository.New(s.config.PostgresConfig, s.logger)

	err := r.Start()
	if err != nil {
		return err
	}

	s.repo = r
	s.logger.Info("Repository configured")
	return nil
}

func (s *Server) startGRPCServer() error {
	if s.repo == nil {
		return errors.New("no repository configured")
	}

	l, err := net.Listen("tcp", ":"+s.config.AppConfig.Port)
	if err != nil {
		s.logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	gRPC.Register(grpcServer, s.repo, s.logger)
	s.logger.Info("Starting gRPCClients Server on port %s", s.config.AppConfig.Port)

	s.grpcServer = grpcServer

	if err := grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}
