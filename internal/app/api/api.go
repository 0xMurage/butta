package api

import (
	"butta/internal/pkg/config"
	"butta/pkg/database"
	"butta/pkg/logger"
	"context"
	"github.com/caarlos0/env/v11"
	"net/http"
)

// Bootstrap includes logic that needs to be executed when api start
func Bootstrap() {}

// Routes returns all registered api routes
func Routes() http.Handler {

	var cfg config.Config

	err := env.Parse(&cfg)
	if err != nil {
		logger.Fatal("unable to parse config", "error", err)
	}

	ctx := context.Background()

	psqlPool, err := database.InitPsqlPool(ctx, cfg.Database.Url)

	if err != nil {
		logger.Fatal("unable to initialize the pgx pool", "error", err)
	}

	return registerRoutes(cfg, psqlPool)
}

// Shutdown performs graceful cleanup of application resources.
func Shutdown() {}
