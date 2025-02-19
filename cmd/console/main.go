package main

import (
	"butta/internal/app/console"
	"butta/internal/pkg/config"
	"butta/pkg/database"
	"butta/pkg/errors"
	"butta/pkg/logger"
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	pool, err := database.InitPsqlPool(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("unable to connect to database pool", "err", err)
	}

	var cfg config.Config

	err = env.Parse(&cfg)
	if err != nil {
		logger.Fatal("unable to parse config", "error", err)
	}

	riverConfig := &river.Config{
		Workers: console.RegisterWorkers(cfg, pool),
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 12},
		},
	}

	jobClient, err := runJobClient(ctx, pool, riverConfig)
	if err != nil {
		logger.Fatal("unable to start job client", "err", err)
	}

	awaitShutdownSignal()

	shutdownJobClient(ctx, jobClient)
}

func runJobClient(ctx context.Context, psqlPool *pgxpool.Pool, config *river.Config) (*river.Client[pgx.Tx], error) {

	consumer, err := river.NewClient(riverpgxv5.New(psqlPool), config)

	if err != nil {
		return nil, errors.Wrap(err, "unable to create job consumer")
	}

	if err = consumer.Start(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to start job consumer")
	}
	return consumer, nil
}

func shutdownJobClient(ctx context.Context, jobClient *river.Client[pgx.Tx]) {
	err := jobClient.Stop(ctx)
	if err != nil {
		logger.Error("Unable to stop job client", "err", err)
	}
}

func awaitShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChan
}
