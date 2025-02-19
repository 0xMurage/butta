package api

import (
	"butta/internal/app/api/middleware"
	"butta/internal/authn"
	"butta/internal/pkg/config"
	"butta/internal/user"
	"butta/pkg/http/router"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func registerRoutes(cfg config.Config, pgPool *pgxpool.Pool) http.Handler {

	mux := router.NewServeMux()

	mux.Use(middleware.Logger) //add middleware that applies to all routes

	authnHandler := authn.New(cfg, pgPool)
	mux.Post("/api/v1/login", authnHandler.Login)

	userHandlers := user.New(cfg, pgPool)
	mux.Get("/api/v1/users", userHandlers.Index)
	mux.Get("/api/v1/users/{id}", userHandlers.Show) //this will return an error for demo purposes
	mux.Post("/api/v1/users", userHandlers.Create)
	mux.Put("/api/v1/users/{id}", userHandlers.Update)
	mux.Delete("/api/v1/users/{id}", userHandlers.Destroy)

	return mux
}
