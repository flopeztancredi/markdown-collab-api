package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
)

func RegisterRoutes(router *gin.Engine, hub *Hub, service *application.Service, originChecker func(r *http.Request) bool) {
	handler := NewHandler(hub, service, originChecker)
	router.GET("/documents/:id/ws", handler.Handle)
}
