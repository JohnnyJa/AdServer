package server

import (
	"context"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/app"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/decisionEngine"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/requests"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/semanticTargetingService"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server interface {
	service.Service
	ConfigureRoute(logger *logrus.Logger, client grpcClients.ProfilesClient, stateClient grpcClients.ProfileStateClient, service semanticTargetingService.SemanticTargetingService) error
}

type server struct {
	config *app.ServerConfig
	gin    *gin.Engine
	logger *logrus.Logger
}

func New(config *app.Config, logger *logrus.Logger) Server {
	r := gin.Default()
	return &server{
		config: config.ServerConfig,
		gin:    r,
		logger: logger,
	}
}

func (r *server) ConfigureRoute(logger *logrus.Logger, client grpcClients.ProfilesClient, stateClient grpcClients.ProfileStateClient, service semanticTargetingService.SemanticTargetingService) error {
	r.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.gin.GET("/dsp", func(c *gin.Context) {
		var bidRequest requests.BidRequest
		if err := c.BindJSON(&bidRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		decisionMaker := decisionEngine.NewDecisionEngine(logger, client, stateClient, service)

		bidResponse, err := decisionMaker.GetWinners(c.Request.Context(), bidRequest)
		if err != nil {
			return
		}

		if bidResponse == nil {
			c.JSON(http.StatusNoContent, gin.H{})
		}

		c.JSON(http.StatusOK, bidResponse)
	})

	return nil
}

func (r *server) Start(ctx context.Context) error {
	if err := r.gin.Run(":" + r.config.Port); err != nil {
		return err
	}
	return nil
}

func (r *server) Stop(ctx context.Context) error {
	return nil
}
