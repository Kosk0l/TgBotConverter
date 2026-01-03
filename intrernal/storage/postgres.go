package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool // пул соединений
}

// Открытие пула соединения
func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error - config_up database: %w", err)
	}

	config.MaxConns = 10 // Максимальное количество открытых соединений
	config.MaxConnLifetime = time.Hour // Максимальное время жизни соединения

	// подключение
	newpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error - create newpool: %w", err)
	}

	return &Postgres{
		pool: newpool,
	}, nil
}

// Закрытие пула соединения
func (p *Postgres) Close() () {
	p.pool.Close()
}

func (p *Postgres) Pool() *pgxpool.Pool {
	return p.pool
}