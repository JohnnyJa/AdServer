package main

import (
	"github.com/JohnnyJa/AdServer/StateService/internal/app"
	"github.com/JohnnyJa/AdServer/StateService/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/StateService/internal/grpcServers"
	"github.com/JohnnyJa/AdServer/StateService/internal/redisStorage"
	"github.com/JohnnyJa/AdServer/StateService/internal/stateManager"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
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
				redisStorage.NewRedisStorage,
				fx.As(new(redisStorage.RedisStorage)),
			),
			fx.Annotate(
				grpcClients.NewProfileLimitClient,
				fx.As(new(grpcClients.ProfilesLimitsClient)),
			),
			fx.Annotate(
				stateManager.NewStateManager,
				fx.As(new(stateManager.StateManager)),
			),
			grpc.NewServer,
		),
		fx.Invoke(
			readConfig,
			configureLogger,
			setupRedis,
			startClient,
			startManager,
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

func setupRedis(redis redisStorage.RedisStorage, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: redis.Start,
		OnStop:  redis.Stop,
	})
	return nil
}

func startClient(client grpcClients.ProfilesLimitsClient, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: client.Start,
		OnStop:  client.Stop,
	})
	return nil
}

func startManager(manager stateManager.StateManager, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: manager.Start,
		OnStop:  manager.Stop,
	})
	return nil
}

func startServer(config *app.Config, logger *logrus.Logger, manager stateManager.StateManager, grpcServer *grpc.Server, lc fx.Lifecycle) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l, err := net.Listen("tcp", ":"+config.Port)
			if err != nil {
				logger.Fatal(err)
			}

			grpcServers.Register(config, logger, manager, grpcServer)
			logger.Info("Starting gRPCClients Server on port %s", config.Port)

			go func() {
				err := grpcServer.Serve(l)
				if err != nil {
					logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})
	return nil
}
