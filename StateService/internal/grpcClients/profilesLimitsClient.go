package grpcClients

import (
	"context"
	"github.com/JohnnyJa/AdServer/StateService/internal/app"
	"github.com/JohnnyJa/AdServer/StateService/internal/grpcClients/proto"
)

type ProfilesLimitsClient interface {
	Client
	GetProfilesWithLimits(ctx context.Context) (profilesLimitsByPackages []*proto.ProfilesWithLimits, err error)
}

type profilesLimitsClient struct {
	*client
	service proto.ProfilesLimitsServiceClient
}

func NewProfileLimitClient(config *app.Config) ProfilesLimitsClient {
	return &profilesLimitsClient{
		client: &client{config: config.ClientConfig},
	}
}

func (c *profilesLimitsClient) Start(ctx context.Context) error {
	err := c.client.Start()
	if err != nil {
		return err
	}
	c.service = proto.NewProfilesLimitsServiceClient(c.conn)
	return nil
}

func (c *profilesLimitsClient) Stop(ctx context.Context) error {
	err := c.client.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (c *profilesLimitsClient) GetProfilesWithLimits(ctx context.Context) (profilesLimitsByPackages []*proto.ProfilesWithLimits, err error) {
	resp, err := c.service.GetProfilesLimits(ctx, &proto.GetProfilesLimitsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.ProfilesWithLimits, nil
}
