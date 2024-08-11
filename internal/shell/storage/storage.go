package storage

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DatabaseURL *url.URL
}

func NewPool(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get new pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to Ping: %w", err)
	}

	return pool, nil
}
