package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab/internal/config"
	healthHTTP "github.com/flopeztancredi/markdown-collab/internal/health/infrastructure/http"
)

func New(cfg *config.Config) *http.Server {
	gin.SetMode(cfg.GinMode)
	router := gin.Default()
	registerRoutes(router, cfg)

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}
}

func registerRoutes(router *gin.Engine, cfg *config.Config) {
	healthHTTP.RegisterRoutes(router, cfg)
}
