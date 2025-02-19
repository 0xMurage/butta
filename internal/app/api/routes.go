package api

import (
	"butta/internal/pkg/config"
	"butta/pkg/http/router"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func registerRoutes(cfg config.Config, pgPool *pgxpool.Pool) http.Handler {

	mux := router.NewServeMux()

	return mux
}
