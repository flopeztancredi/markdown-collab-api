package http

import (
	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
)

func RegisterRoutes(router *gin.Engine, service *application.Service) {
	handler := NewHandler(service)

	router.POST("/documents", handler.Create)
	router.GET("/documents/:id", handler.GetByID)
}
