package router

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
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
		var bytes []byte
		if c.Request.Body != nil {
			bytes, _ = io.ReadAll(c.Request.Body)
		}

		r.workers.Write(string(bytes))
		c.JSON(http.StatusOK, gin.H{})
	})
}

func (r *Router) Run(port string) error {
	if err := r.router.Run(port); err != nil {
		return err
	}
	return nil
}
