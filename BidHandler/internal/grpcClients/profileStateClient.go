package grpcClients

import (
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients/proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type ProfileStateClient interface {
	Client
	GetProfileState(ctx context.Context, profileId uuid.UUID) (int, error)
}

type profileStateClient struct {
	client
	service proto.ProfileStateServiceClient
}

func NewProfileStateClient(config *app.Config) ProfileStateClient {
	return &profileStateClient{
		client: client{config: config.ProfilesStateClientConfig},
	}
}

func (c *profileStateClient) Start(ctx context.Context) error {
	err := c.client.Start()
	if err != nil {
		return err
	}
	c.service = proto.NewProfileStateServiceClient(c.conn)
	return nil
}

func (c *profileStateClient) Stop(ctx context.Context) error {
	err := c.client.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (c *profileStateClient) GetProfileState(ctx context.Context, profileId uuid.UUID) (int, error) {
	resp, err := c.service.GetProfileState(ctx, &proto.GetProfileStateRequest{ProfileId: profileId.String()})
	if err != nil {
		return 0, err
	}

	return int(resp.State), nil
}
