package mysql

import (
	"Advertisement/internal/domain"
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
)

const adstable = "advertisement"

var tracer = otel.GetTracerProvider().Tracer("advertisement")

func (r *AdvertisementMysql) CreateAdvertisement(ctx context.Context, ad *domain.Advertisement) error {
	ctx, span := tracer.Start(ctx, "Create-mysql")
	defer span.End()
	query := fmt.Sprintf("INSERT INTO %s (title, description, price) VALUES (?, ?, ?)", adstable)
	_, err := r.db.ExecContext(ctx, query, ad.Title, ad.Description, ad.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdvertisementMysql) GetAdvertisementByID(ctx context.Context, id int) (*domain.Advertisement, error) {
	ctx, span := tracer.Start(ctx, "GetById-mysql")
	defer span.End()
	var ad domain.Advertisement
	err := r.db.QueryRowContext(ctx, `
        SELECT id, title, description, price, createdTime, isActive
        FROM advertisements
        WHERE id = ?
    `, id).Scan(
		&ad.Id,
		&ad.Title,
		&ad.Description,
		&ad.Price,
		&ad.CreatedTime,
		&ad.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &ad, nil
}

func (r *AdvertisementMysql) GetAllSortedAndPaginated(ctx context.Context, sortBy, sortOrder string, offset, limit int) ([]*domain.Advertisement, error) {
	ctx, span := tracer.Start(ctx, "GetAll-mysql")
	defer span.End()
	query := fmt.Sprintf("SELECT id,title,description,price,createdTime,isActive FROM advertisements ORDER BY %s %s LIMIT ?,?", sortBy, sortOrder)
	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	advertisements := []*domain.Advertisement{}

	for rows.Next() {
		var ad domain.Advertisement
		if err := rows.Scan(
			&ad.Id,
			&ad.Title,
			&ad.Description,
			&ad.Price,
			&ad.CreatedTime,
			&ad.IsActive,
		); err != nil {
			return nil, err
		}
		advertisements = append(advertisements, &ad)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (r *AdvertisementMysql) UpdateAdvertisement(ctx context.Context, id int, ad *domain.UpdateAdvertisementInput) error {
	ctx, span := tracer.Start(ctx, "Update-mysql")
	defer span.End()
	query := `UPDATE advertisements SET title=?, description=?, price=?, isActive=? WHERE id=?`

	_, err := r.db.ExecContext(ctx, query, ad.Title, ad.Description, ad.Price, ad.IsActive, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdvertisementMysql) DeleteAdvertisement(ctx context.Context, id int) error {
	ctx, span := tracer.Start(ctx, "Delete-mysql")
	defer span.End()
	query := `DELETE FROM advertisements WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
