package main

import (
	"context"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/server"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
)

func main() {
	fx.New(CreateApp()).Run()
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			app.NewConfig,
			logrus.New,
			fx.Annotate(
				grpcClients.NewProfileClient,
				fx.As(new(grpcClients.ProfilesClient)),
			),
			fx.Annotate(
				server.New,
				fx.As(new(server.Server)),
			),
		),
		fx.Invoke(
			readConfig,
			configureLogger,
			setupClient,
			startServer,
		),
	)

}

func readConfig(config *app.Config) error {
	err := config.ReadConfig()
	if err != nil {
		return err
	}
	return nil
}

func configureLogger(config *app.Config, logger *logrus.Logger) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logger.Info("Logger configured")

	return nil
}

func setupClient(client grpcClients.ProfilesClient, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: client.Start,
		OnStop:  client.Stop,
	})
	return nil
}

func startServer(srv server.Server, client grpcClients.ProfilesClient, logger *logrus.Logger, lc fx.Lifecycle) error {
	err := srv.ConfigureRoute(logger, client)
	if err != nil {
		return err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.Start(ctx)
				if err != nil {
					logger.Errorf("Error starting server: %v", err)
				}
			}()
			return nil
		},
		OnStop: srv.Stop,
	})
	return nil
}
