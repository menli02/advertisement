package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/models"
)

type ImageRepository struct {
	pool *pgxpool.Pool
}

func NewImageRepository(pool *pgxpool.Pool) *ImageRepository {
	return &ImageRepository{pool: pool}
}

func (r *ImageRepository) ReplaceAll(ctx context.Context, adID int64, urls []string) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `DELETE FROM advertisement_images WHERE ad_id = $1`, adID); err != nil {
		return err
	}

	for i, url := range urls {
		if _, err := tx.Exec(ctx,
			`INSERT INTO advertisement_images (ad_id, url, position) VALUES ($1, $2, $3)`,
			adID, url, i,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *ImageRepository) GetByAdID(ctx context.Context, adID int64) ([]*models.AdImage, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, ad_id, url, position, created_at FROM advertisement_images
		 WHERE ad_id = $1 ORDER BY position ASC`,
		adID,
	)
	if err != nil {
		return nil, err
	}
	imgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.AdImage])
	if err != nil {
		return nil, err
	}
	result := make([]*models.AdImage, len(imgs))
	for i, img := range imgs {
		img := img
		result[i] = &img
	}
	return result, nil
}
