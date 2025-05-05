package gRPC

import (
	"context"
	"github.com/JohnnyJa/AdServer/PackageService/internal/gRPC/proto"
	"github.com/JohnnyJa/AdServer/PackageService/internal/mapper"
	"github.com/JohnnyJa/AdServer/PackageService/internal/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedPackageServiceServer
	repo   repository.Repository
	logger *logrus.Logger
}

func Register(gRPC *grpc.Server, repo repository.Repository, logger *logrus.Logger) {
	proto.RegisterPackageServiceServer(gRPC, &server{repo: repo, logger: logger})
}

func (s *server) GetPackagesWithZones(ctx context.Context, req *proto.GetPackagesWithZonesRequest) (*proto.GetPackagesWithZonesResponse, error) {

	ids, err := mapper.StringsToUUIDs(req.PackageIds)
	if err != nil {
		return nil, err
	}

	packages, err := s.repo.ReadPackages(ctx, ids)

	if err != nil {
		return nil, err
	}

	var result []*proto.Package
	for _, pkg := range packages {
		result = append(result, &proto.Package{
			Id:      pkg.Id.String(),
			Name:    pkg.Name,
			ZoneIds: mapper.UUIDsToStrings(pkg.ZoneIds),
		})
	}

	return &proto.GetPackagesWithZonesResponse{Packages: result}, nil
}
