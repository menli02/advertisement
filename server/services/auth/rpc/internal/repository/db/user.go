package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/models"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, phone, first_name, last_name, created_at, updated_at
		 FROM users WHERE phone = $1 LIMIT 1`,
		phone,
	)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Upsert(ctx context.Context, phone, firstName, lastName string) (*models.User, error) {
	rows, err := r.pool.Query(ctx,
		`INSERT INTO users (phone, first_name, last_name)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (phone) DO UPDATE
		   SET first_name = EXCLUDED.first_name,
		       last_name  = EXCLUDED.last_name,
		       updated_at = NOW()
		 RETURNING id, phone, first_name, last_name, created_at, updated_at`,
		phone, firstName, lastName,
	)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}
	return &user, nil
}
