package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct{}

func New() *gin.Engine {
	r := gin.Default()
	configureRouter(r)
	return r
}

func configureRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
