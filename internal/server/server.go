package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flopeztancredi/markdown-collab-api/internal/config"
	documentHTTP "github.com/flopeztancredi/markdown-collab-api/internal/document/infrastructure/http"
	healthHTTP "github.com/flopeztancredi/markdown-collab-api/internal/health/infrastructure/http"
)

func New(cfg *config.Config, db *pgxpool.Pool) *http.Server {
	gin.SetMode(cfg.GinMode)
	router := gin.Default()
	registerRoutes(router, db)

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}
}

func registerRoutes(router *gin.Engine, db *pgxpool.Pool) {
	healthHTTP.RegisterRoutes(router)
	documentHTTP.RegisterRoutes(router, db)
}
