package gRPCClients

import (
	"context"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCClients/proto"
)

type ProfileClient interface {
	Client
	GetProfiles() ([]*proto.Profile, error)
}

type profileClient struct {
	client
	service proto.ProfilesServiceClient
}

func NewProfileClient(config *Config) ProfileClient {
	return &profileClient{
		client: client{
			config: config,
		},
	}
}

func (c *profileClient) Start() error {
	err := c.client.Start()
	if err != nil {
		return err
	}
	c.service = proto.NewProfilesServiceClient(c.conn)
	return nil
}

func (c *profileClient) GetProfiles() ([]*proto.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.config.Timeout)
	defer cancel()
	resp, err := c.service.GetActiveProfiles(ctx, &proto.GetProfilesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Profiles, nil
}
