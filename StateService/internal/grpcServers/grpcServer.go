package grpcServers

import (
	"github.com/JohnnyJa/AdServer/StateService/internal/app"
	"github.com/JohnnyJa/AdServer/StateService/internal/grpcServers/proto"
	"github.com/JohnnyJa/AdServer/StateService/internal/stateManager"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedUpdateCounterServiceServer
	proto.UnimplementedProfileStateServiceServer
	conf    *app.ServerConfig
	logger  *logrus.Logger
	manager stateManager.StateManager
}

func Register(conf *app.Config, logger *logrus.Logger, manager stateManager.StateManager, grpc *grpc.Server) {
	proto.RegisterProfileStateServiceServer(grpc, &server{conf: conf.ServerConfig, logger: logger, manager: manager})
	proto.RegisterUpdateCounterServiceServer(grpc, &server{conf: conf.ServerConfig, logger: logger, manager: manager})
}

func (s *server) UpdateCounterOnProfile(ctx context.Context, in *proto.UpdateCounterRequest) (*proto.UpdateCounterResponse, error) {
	err := s.manager.IncrementViews(ctx, uuid.MustParse(in.ProfileId))
	if err != nil {
		return nil, err
	}
	return &proto.UpdateCounterResponse{}, nil
}

func (s *server) GetProfileState(ctx context.Context, in *proto.GetProfileStateRequest) (*proto.GetProfileStateResponse, error) {
	state, err := s.manager.GetState(ctx, uuid.MustParse(in.ProfileId))
	if err != nil {
		return nil, err
	}
	return &proto.GetProfileStateResponse{State: int32(state)}, nil
}
