package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(dsn string) (*Postgres, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &Postgres{pool: pool}, nil
}

func (p *Postgres) CreateSession(ctx context.Context, id, userID string) error {
	_, err := p.pool.Exec(ctx,
		`INSERT INTO sessions (id, user_id, active) VALUES ($1, $2, true)`, id, userID)
	return err
}
