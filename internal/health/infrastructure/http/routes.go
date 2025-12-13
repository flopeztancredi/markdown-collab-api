package http

import (
	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab/internal/config"
	"github.com/flopeztancredi/markdown-collab/internal/health/application"
)

func RegisterRoutes(router *gin.Engine, cfg *config.Config) {
	service := application.NewService(cfg.AppName, cfg.AppVersion)
	handler := NewHandler(service)
	router.GET("/health", handler.Check)
}
