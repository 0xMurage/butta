package app

import (
	"butta/internal/authn"
	"butta/internal/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
)

func RegisterWorkers(cfg config.Config, psqlPool *pgxpool.Pool) *river.Workers {
	workers := river.NewWorkers()

	river.AddWorker(workers, &authn.SendPasswordResetLinkWorker{PsqlPool: psqlPool})

	return workers
}
