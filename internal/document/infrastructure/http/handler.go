package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{service: service}
}

type CreateRequest struct {
	Title string `json:"title"`
}

type DocumentResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	title := req.Title
	if title == "" {
		title = "Untitled"
	}

	doc, err := h.service.Create(c.Request.Context(), title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create document"})
		return
	}

	c.JSON(http.StatusCreated, DocumentResponse{
		ID:        doc.ID.String(),
		Title:     doc.Title,
		CreatedAt: doc.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: doc.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document id"})
		return
	}

	doc, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		return
	}

	c.JSON(http.StatusOK, DocumentResponse{
		ID:        doc.ID.String(),
		Title:     doc.Title,
		CreatedAt: doc.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: doc.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}
