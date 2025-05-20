package gRPC

import (
	"context"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/gRPC/proto"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/mapper"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedProfilesServiceServer
	repo   repository.Repository
	logger *logrus.Logger
}

func Register(gRPC *grpc.Server, repo repository.Repository, logger *logrus.Logger) {
	proto.RegisterProfilesServiceServer(gRPC, &server{repo: repo, logger: logger})
}

func (s *server) GetActiveProfiles(ctx context.Context, in *proto.GetProfilesRequest) (*proto.GetProfilesResponse, error) {
	s.logger.Infof("GetActiveProfiles called with args %+v", in)
	profiles, err := s.repo.ReadProfiles(ctx)
	if err != nil {
		return nil, err
	}

	toProfiles, err := mapper.ToProfiles(profiles)
	if err != nil {
		return nil, err
	}

	return mapper.ToGrpcProfiles(toProfiles), nil
}

func (s *server) GetProfilesLimits(ctx context.Context, in *proto.GetProfilesLimitsRequest) (*proto.GetProfilesLimitsResponse, error) {
	s.logger.Infof("GetProfilesLimits called with args %+v", in)
	profiles, err := s.repo.ReadProfilesLimits(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	grpcProfiles := mapper.ToGrpcProfilesLimits(profiles)
	return grpcProfiles, nil
}
