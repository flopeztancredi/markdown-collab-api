package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flopeztancredi/markdown-collab-api/internal/health/application"
)

type Handler struct {
	service *application.Service
}

func NewHandler() *Handler {
	return &Handler{service: application.NewService()}
}

func (h *Handler) Check(c *gin.Context) {
	status := h.service.Check()
	c.JSON(http.StatusOK, status)
}
