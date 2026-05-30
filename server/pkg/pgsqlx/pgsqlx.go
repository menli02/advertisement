package pgsqlx

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	DSN string
}

type DB struct {
	pool *pgxpool.Pool
}

func NewPostgresSQL(cfg Config) *DB {
	pool, err := pgxpool.New(context.Background(), cfg.DSN)
	if err != nil {
		logx.Errorf("pgsqlx: failed to create pool: %v", err)
		panic(err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		logx.Errorf("pgsqlx: failed to ping: %v", err)
		panic(err)
	}
	return &DB{pool: pool}
}

func (d *DB) Pool() *pgxpool.Pool {
	return d.pool
}
