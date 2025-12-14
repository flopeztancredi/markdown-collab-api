package server

import (
	"context"
	"net/http"
	"net/url"
	"time"

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
	router.Use(requestTimeoutMiddleware(cfg.RequestTimeout))

	documentRepo := documentPostgres.NewRepository(db)
	documentService := application.NewService(documentRepo, cfg.DBTimeout)
	documentHub := documentWS.NewHub()
	go documentHub.Run()

	healthHTTP.RegisterRoutes(router)
	documentHTTP.RegisterRoutes(router, documentService)
	documentWS.RegisterRoutes(router, documentHub, documentService, allowedOriginChecker(cfg.AllowedOrigins))

	return &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}
}

func allowedOriginChecker(origins []string) func(r *http.Request) bool {
	allowed := make(map[string]struct{}, len(origins))
	for _, o := range origins {
		allowed[o] = struct{}{}
	}

	return func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return false
		}

		parsed, err := url.Parse(origin)
		if err != nil {
			return false
		}

		if isLoopbackHost(parsed.Hostname()) {
			return true
		}

		if len(allowed) == 0 {
			return true
		}

		_, ok := allowed[origin]
		return ok
	}
}

func isLoopbackHost(host string) bool {
	switch host {
	case "localhost", "127.0.0.1", "::1":
		return true
	default:
		return false
	}
}

func requestTimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	if timeout <= 0 {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
