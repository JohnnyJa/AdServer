package main

import (
	"context"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/semanticTargetingService"
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
				semanticTargetingService.New,
				fx.As(new(semanticTargetingService.SemanticTargetingService))),
			fx.Annotate(
				grpcClients.NewProfileClient,
				fx.As(new(grpcClients.ProfilesClient)),
			),
			fx.Annotate(
				grpcClients.NewProfileStateClient,
				fx.As(new(grpcClients.ProfileStateClient)),
			),
			fx.Annotate(
				server.New,
				fx.As(new(server.Server)),
			),
		),
		fx.Invoke(
			readConfig,
			configureLogger,
			startSemanticService,
			setupProfilesClient,
			setupProfileStateClient,
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

func startSemanticService(service semanticTargetingService.SemanticTargetingService, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: service.Start,
		OnStop:  service.Stop,
	})
	return nil
}

func setupProfilesClient(client grpcClients.ProfilesClient, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: client.Start,
		OnStop:  client.Stop,
	})
	return nil
}

func setupProfileStateClient(client grpcClients.ProfileStateClient, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: client.Start,
		OnStop:  client.Stop,
	})
	return nil
}

func startServer(srv server.Server, profilesClient grpcClients.ProfilesClient, profileStateClient grpcClients.ProfileStateClient, service semanticTargetingService.SemanticTargetingService, logger *logrus.Logger, lc fx.Lifecycle) error {
	err := srv.ConfigureRoute(logger, profilesClient, profileStateClient, service)
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
