package http

import (
	"go-server/internal/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(cfg *config.Configs) (*Router, error) {
	//create gin instance
	router := gin.New()
	//create handlers
	handler := newHandler()

	//config gin logger
	if cfg.WebServer.GIN.UseLogger {
		router.Use(gin.Logger())
	}

	//config routes
	rg := router.Group("/api")
	{
		match := rg.Group("/help")
		{
			match.GET("/get", handler.help.Get)
		}
	}

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	return &Router{router}, nil
}
