package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab-api/internal/config"
	healthHTTP "github.com/flopeztancredi/markdown-collab-api/internal/health/infrastructure/http"
)

func New(cfg *config.Config) *http.Server {
	gin.SetMode(cfg.GinMode)
	router := gin.Default()
	registerRoutes(router)

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}
}

func registerRoutes(router *gin.Engine) {
	healthHTTP.RegisterRoutes(router)
}
