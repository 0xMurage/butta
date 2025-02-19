package api

import (
	"butta/internal/app/api/middleware"
	"butta/internal/pkg/config"
	"butta/pkg/http/router"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func registerRoutes(cfg config.Config, pgPool *pgxpool.Pool) http.Handler {

	mux := router.NewServeMux()

	mux.Use(middleware.Logger) //add middleware that applies to all routes

	return mux
}
