package grpcClients

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/grpcClients/proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type IncrementViewsClient interface {
	Client
	IncrementViews(ctx context.Context, profileId uuid.UUID) error
}

type incrementViewsClient struct {
	*client
	service proto.UpdateCounterServiceClient
}

func NewIncrementViewsClient(config *ClientConfig) IncrementViewsClient {
	return &incrementViewsClient{
		client: &client{config: config},
	}
}

func (c *incrementViewsClient) Start() error {
	err := c.client.Start()
	if err != nil {
		return err
	}
	c.service = proto.NewUpdateCounterServiceClient(c.conn)
	return nil
}

func (c *incrementViewsClient) Stop() error {
	err := c.client.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (c *incrementViewsClient) IncrementViews(ctx context.Context, profileId uuid.UUID) error {
	_, err := c.service.UpdateCounterOnProfile(ctx, &proto.UpdateCounterRequest{ProfileId: profileId.String()})
	if err != nil {
		return err
	}
	return nil
}
