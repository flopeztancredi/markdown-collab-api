package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
	"github.com/flopeztancredi/markdown-collab-api/internal/document/infrastructure/postgres"
)

func RegisterRoutes(router *gin.Engine, db *pgxpool.Pool) {
	repo := postgres.NewRepository(db)
	service := application.NewService(repo)
	handler := NewHandler(service)

	router.POST("/documents", handler.Create)
	router.GET("/documents/:id", handler.GetByID)
}
