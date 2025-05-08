package grpcClients

import (
	"context"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients/proto"
	"github.com/google/uuid"
)

type ProfilesClient interface {
	Client
	GetProfilesMapsByZone(ctx context.Context, zoneId uuid.UUID) (profilesByPackages map[string]*proto.ProfileIds, profilesByUUID map[string]*proto.Profile, err error)
}

type profilesClient struct {
	client
	service proto.ProfilesByZoneServiceClient
}

func NewProfileClient(config *app.Config) ProfilesClient {
	return &profilesClient{
		client: client{config: config.ClientConfig},
	}
}

func (c *profilesClient) Start(ctx context.Context) error {
	err := c.client.Start()
	if err != nil {
		return err
	}
	c.service = proto.NewProfilesByZoneServiceClient(c.conn)
	return nil
}

func (c *profilesClient) Stop(ctx context.Context) error {
	err := c.client.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (c *profilesClient) GetProfilesMapsByZone(ctx context.Context, zoneId uuid.UUID) (profilesByPackages map[string]*proto.ProfileIds, profilesByUUID map[string]*proto.Profile, err error) {
	resp, err := c.service.GetProfilesByZone(ctx, &proto.GetProfileByZoneRequest{ZoneId: zoneId.String()})
	if err != nil {
		return nil, nil, err
	}

	profilesByPackages = resp.ProfilesByPackage
	profilesByUUID = resp.ProfilesByUUID
	return
}
