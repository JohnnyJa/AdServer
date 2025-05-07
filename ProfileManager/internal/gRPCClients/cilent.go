package gRPCClients

import (
	"github.com/JohnnyJa/AdServer/ProfileManager/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	service.Service
}

type client struct {
	config *Config
	conn   *grpc.ClientConn
}

func NewClient(config *Config) Client {
	return &client{
		config: config,
	}
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
