package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPsqlPool(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	psqlPool, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		return nil, err
	}

	conn, err := psqlPool.Acquire(ctx)

	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)

	if err != nil {
		return nil, err
	}
	return psqlPool, nil
}
