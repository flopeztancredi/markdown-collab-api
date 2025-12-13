package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	handler := NewHandler()
	router.GET("/health", handler.Check)
}
