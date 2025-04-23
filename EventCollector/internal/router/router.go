package router

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/model"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Router struct {
	router  *gin.Engine
	logger  *logrus.Logger
	workers *worker.Pool
}

func New(workers *worker.Pool, logger *logrus.Logger) *Router {
	r := gin.Default()
	return &Router{
		router:  r,
		logger:  logger,
		workers: workers,
	}
}

func (r *Router) Start() {
	r.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.router.GET("/events", func(c *gin.Context) {
		var event model.Event
		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		r.workers.Write(event)
		c.JSON(http.StatusOK, gin.H{})
	})
}

func (r *Router) Run(port string) error {
	if err := r.router.Run(port); err != nil {
		return err
	}
	return nil
}
