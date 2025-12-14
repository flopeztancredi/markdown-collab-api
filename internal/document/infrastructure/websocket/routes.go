package websocket

import (
	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
)

func RegisterRoutes(router *gin.Engine, hub *Hub, service *application.Service) {
	handler := NewHandler(hub, service)
	router.GET("/documents/:id/ws", handler.Handle)
}
