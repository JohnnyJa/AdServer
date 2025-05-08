package grpcClients

import (
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	service.Service
}

type client struct {
	config *app.ClientConfig
	conn   *grpc.ClientConn
}

func (c *client) Start() error {
	conn, err := grpc.NewClient(c.config.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) Stop() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}

	return nil
}
