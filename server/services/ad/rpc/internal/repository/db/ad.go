package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/models"
)

var allowedSortColumns = map[string]bool{
	"created_at": true,
	"updated_at": true,
	"price":      true,
	"view_count": true,
	"title":      true,
}

type AdRepository struct {
	pool *pgxpool.Pool
}

func NewAdRepository(pool *pgxpool.Pool) *AdRepository {
	return &AdRepository{pool: pool}
}

func (r *AdRepository) Create(ctx context.Context, ad *models.Advertisement) (*models.Advertisement, error) {
	rows, err := r.pool.Query(ctx,
		`INSERT INTO advertisements (user_id, category_id, title, description, slug, price, currency, status)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, 'active')
		 RETURNING id, user_id, category_id, title, description, slug, price, currency, status, view_count, created_at, updated_at, deleted_at`,
		ad.UserID, ad.CategoryID, ad.Title, ad.Description, ad.Slug, ad.Price, ad.Currency,
	)
	if err != nil {
		return nil, err
	}
	created, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Advertisement])
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (r *AdRepository) GetByID(ctx context.Context, id int64) (*models.Advertisement, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, category_id, title, description, slug, price, currency, status, view_count, created_at, updated_at, deleted_at
		 FROM advertisements WHERE id = $1 AND deleted_at IS NULL LIMIT 1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	ad, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Advertisement])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ad, nil
}

func (r *AdRepository) GetBySlug(ctx context.Context, slug string) (*models.Advertisement, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, category_id, title, description, slug, price, currency, status, view_count, created_at, updated_at, deleted_at
		 FROM advertisements WHERE slug = $1 AND deleted_at IS NULL LIMIT 1`,
		slug,
	)
	if err != nil {
		return nil, err
	}
	ad, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Advertisement])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &ad, nil
}

func (r *AdRepository) List(ctx context.Context, categoryID, userID int64, query, sortBy, sortOrder string, page, pageSize int32) ([]*models.Advertisement, int64, error) {
	// Whitelist sort column to prevent SQL injection
	col := "created_at"
	if allowedSortColumns[sortBy] {
		col = sortBy
	}
	dir := "DESC"
	if strings.ToUpper(sortOrder) == "ASC" {
		dir = "ASC"
	}

	args := []interface{}{}
	filters := []string{"deleted_at IS NULL"}
	i := 1

	if categoryID > 0 {
		filters = append(filters, fmt.Sprintf("category_id = $%d", i))
		args = append(args, categoryID)
		i++
	}
	if userID > 0 {
		filters = append(filters, fmt.Sprintf("user_id = $%d", i))
		args = append(args, userID)
		i++
	}
	if query != "" {
		filters = append(filters, fmt.Sprintf("(title ILIKE $%d OR description ILIKE $%d)", i, i))
		args = append(args, "%"+query+"%")
		i++
	}

	where := strings.Join(filters, " AND ")

	var total int64
	countRow := r.pool.QueryRow(ctx, fmt.Sprintf("SELECT COUNT(*) FROM advertisements WHERE %s", where), args...)
	if err := countRow.Scan(&total); err != nil {
		return nil, 0, err
	}

	if pageSize <= 0 {
		pageSize = 20
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	listArgs := append(args, pageSize, offset)
	sql := fmt.Sprintf(
		`SELECT id, user_id, category_id, title, description, slug, price, currency, status, view_count, created_at, updated_at, deleted_at
		 FROM advertisements WHERE %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
		where, col, dir, i, i+1,
	)
	rows, err := r.pool.Query(ctx, sql, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	ads, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Advertisement])
	if err != nil {
		return nil, 0, err
	}
	result := make([]*models.Advertisement, len(ads))
	for i, a := range ads {
		a := a
		result[i] = &a
	}
	return result, total, nil
}

func (r *AdRepository) Update(ctx context.Context, ad *models.Advertisement) (*models.Advertisement, error) {
	rows, err := r.pool.Query(ctx,
		`UPDATE advertisements
		 SET category_id = $1, title = $2, description = $3, price = $4, currency = $5, status = $6, updated_at = NOW()
		 WHERE id = $7 AND user_id = $8 AND deleted_at IS NULL
		 RETURNING id, user_id, category_id, title, description, slug, price, currency, status, view_count, created_at, updated_at, deleted_at`,
		ad.CategoryID, ad.Title, ad.Description, ad.Price, ad.Currency, ad.Status, ad.ID, ad.UserID,
	)
	if err != nil {
		return nil, err
	}
	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Advertisement])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &updated, nil
}

func (r *AdRepository) SoftDelete(ctx context.Context, id, userID int64) (bool, error) {
	tag, err := r.pool.Exec(ctx,
		`UPDATE advertisements SET deleted_at = NOW() WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL`,
		id, userID,
	)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() > 0, nil
}

func (r *AdRepository) IncrementView(ctx context.Context, id int64) (int64, error) {
	var count int64
	err := r.pool.QueryRow(ctx,
		`UPDATE advertisements SET view_count = view_count + 1 WHERE id = $1 AND deleted_at IS NULL RETURNING view_count`,
		id,
	).Scan(&count)
	return count, err
}

func (r *AdRepository) SlugExists(ctx context.Context, slug string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM advertisements WHERE slug = $1 AND deleted_at IS NULL)`,
		slug,
	).Scan(&exists)
	return exists, err
}
