package main

import (
	"github.com/JohnnyJa/AdServer/StateService/internal/app"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
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
