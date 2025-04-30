package router

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/kafka"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/model"
	"github.com/JohnnyJa/AdServer/EventCollector/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Router interface {
	service.Service
	ConfigureRoute() error
}

type router struct {
	port   string
	gin    *gin.Engine
	logger *logrus.Logger
	kafka  kafka.Kafka
}

func New(port string, logger *logrus.Logger, kafka kafka.Kafka) Router {
	r := gin.Default()
	return &router{
		port:   port,
		gin:    r,
		logger: logger,
		kafka:  kafka,
	}
}

func (r *router) ConfigureRoute() error {
	r.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.gin.GET("/events", func(c *gin.Context) {
		var event model.Event
		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		err := r.kafka.Write(event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	return nil
}

func (r *router) Start() error {
	if err := r.gin.Run(r.port); err != nil {
		return err
	}
	return nil
}

func (r *router) Stop() error {
	return nil
}
