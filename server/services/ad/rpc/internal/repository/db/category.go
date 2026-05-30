package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/models"
)

type CategoryRepository struct {
	pool *pgxpool.Pool
}

func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{pool: pool}
}

func (r *CategoryRepository) Create(ctx context.Context, name, slug string) (*models.Category, error) {
	rows, err := r.pool.Query(ctx,
		`INSERT INTO categories (name, slug) VALUES ($1, $2)
		 ON CONFLICT (slug) DO UPDATE SET name = EXCLUDED.name
		 RETURNING id, name, slug, created_at`,
		name, slug,
	)
	if err != nil {
		return nil, err
	}
	cat, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int64) (*models.Category, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, name, slug, created_at FROM categories WHERE id = $1 LIMIT 1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	cat, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cat, nil
}

func (r *CategoryRepository) List(ctx context.Context) ([]*models.Category, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, name, slug, created_at FROM categories ORDER BY name ASC`,
	)
	if err != nil {
		return nil, err
	}
	cats, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}
	result := make([]*models.Category, len(cats))
	for i, c := range cats {
		c := c
		result[i] = &c
	}
	return result, nil
}
