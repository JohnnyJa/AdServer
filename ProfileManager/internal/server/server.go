package server

import (
	"errors"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCClients"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCServer"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/profileStorage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

type Server struct {
	config        *Config
	logger        *logrus.Logger
	profileClient gRPCClients.ProfileClient
	packageClient gRPCClients.PackageClient
	grpcServer    *grpc.Server
	storage       profileStorage.ProfileStorage
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

	if err := s.configureClients(); err != nil {
		return err
	}

	if err := s.configureStorage(); err != nil {
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

func (s *Server) configureStorage() error {
	if s.packageClient == nil {
		return errors.New("no package client configured")
	}

	if s.profileClient == nil {
		return errors.New("no profile client configured")
	}

	storage := profileStorage.New(s.config.StorageConfig, s.logger, s.profileClient, s.packageClient)

	err := storage.Start()
	if err != nil {
		return err
	}

	s.storage = storage

	s.logger.Info("Storage configured")

	return nil
}

func (s *Server) startGRPCServer() error {
	if s.storage == nil {
		return errors.New("no storage configured")
	}

	l, err := net.Listen("tcp", ":"+s.config.AppConfig.Port)
	if err != nil {
		s.logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	gRPCServer.Register(grpcServer, s.storage, s.logger)
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

func (s *Server) configureClients() error {
	var err error
	s.profileClient, err = s.configureProfileClient()
	if err != nil {
		return err
	}
	s.logger.Info("Configured profile client")

	s.packageClient, err = s.configurePackageClient()
	if err != nil {
		return err
	}
	s.logger.Info("Configured package client")
	return nil
}

func (s *Server) configureProfileClient() (gRPCClients.ProfileClient, error) {
	cl := gRPCClients.NewProfileClient(s.config.ProfileClientConfig)
	err := cl.Start()
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (s *Server) configurePackageClient() (gRPCClients.PackageClient, error) {
	cl := gRPCClients.NewPackageClient(s.config.PackageClientConfig)
	err := cl.Start()
	if err != nil {
		return nil, err
	}

	return cl, nil
}
