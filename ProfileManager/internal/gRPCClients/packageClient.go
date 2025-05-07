package gRPCClients

import (
	"context"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCClients/proto"
)

type PackageClient interface {
	Client
	GetPackages(PackageIds []string) ([]*proto.Package, error)
}

type packageClient struct {
	client
	service proto.PackageServiceClient
}

func NewPackageClient(config *Config) PackageClient {
	return &packageClient{
		client: client{
			config: config,
		},
	}
}

func (c *packageClient) Start() error {
	err := c.client.Start()
	if err != nil {
		return err
	}
	c.service = proto.NewPackageServiceClient(c.conn)
	return nil
}

func (c *packageClient) GetPackages(PackageIds []string) ([]*proto.Package, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
	defer cancel()
	resp, err := c.service.GetPackagesWithZones(ctx, &proto.GetPackagesWithZonesRequest{
		PackageIds: PackageIds,
	})
	if err != nil {
		return nil, err
	}
	return resp.Packages, nil
}
