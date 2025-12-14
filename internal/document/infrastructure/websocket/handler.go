package websocket

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gorillaws "github.com/gorilla/websocket"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
)

type Handler struct {
	hub      *Hub
	service  *application.Service
	upgrader gorillaws.Upgrader
}

func NewHandler(hub *Hub, service *application.Service, originChecker func(r *http.Request) bool) *Handler {
	return &Handler{
		hub:     hub,
		service: service,
		upgrader: gorillaws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     originChecker,
		},
	}
}

func (h *Handler) Handle(c *gin.Context) {
	documentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document id"})
		return
	}

	if _, err := h.service.GetByID(c.Request.Context(), documentID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}

	client := NewClient(h.hub, conn, documentID)
	h.hub.Register() <- client

	go client.WritePump()
	client.ReadPump(func(message []byte) {
		go h.persistDocumentUpdate(documentID, message)
	})
}

func (h *Handler) persistDocumentUpdate(documentID uuid.UUID, content []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.UpdateContent(ctx, documentID, content); err != nil {
		log.Printf("failed to persist document %s: %v", documentID, err)
	}
}
