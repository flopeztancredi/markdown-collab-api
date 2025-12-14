package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flopeztancredi/markdown-collab-api/internal/config"
	"github.com/flopeztancredi/markdown-collab-api/internal/document/application"
	documentHTTP "github.com/flopeztancredi/markdown-collab-api/internal/document/infrastructure/http"
	documentPostgres "github.com/flopeztancredi/markdown-collab-api/internal/document/infrastructure/postgres"
	documentWS "github.com/flopeztancredi/markdown-collab-api/internal/document/infrastructure/websocket"
	healthHTTP "github.com/flopeztancredi/markdown-collab-api/internal/health/infrastructure/http"
)

func New(cfg *config.Config, db *pgxpool.Pool) *http.Server {
	gin.SetMode(cfg.GinMode)
	router := gin.Default()

	documentRepo := documentPostgres.NewRepository(db)
	documentService := application.NewService(documentRepo)
	documentHub := documentWS.NewHub()
	go documentHub.Run()

	healthHTTP.RegisterRoutes(router)
	documentHTTP.RegisterRoutes(router, documentService)
	documentWS.RegisterRoutes(router, documentHub, documentService)

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}
}
