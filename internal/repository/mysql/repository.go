package mysql

import (
	"Advertisement/internal/domain"
	"context"
	"database/sql"
)

type Repository interface {
	CreateAdvertisement(ctx context.Context, ad *domain.Advertisement) error
	GetAdvertisementByID(ctx context.Context, id int) (*domain.Advertisement, error)
	GetAllSortedAndPaginated(ctx context.Context, sortBy, sortOrder string, offset, limit int) ([]*domain.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, id int, ad *domain.UpdateAdvertisementInput) error
	DeleteAdvertisement(ctx context.Context, id int) error
}
type AdvertisementMysql struct {
	db *sql.DB
}

func NewAdvertisementMysql(db *sql.DB) *AdvertisementMysql {
	return &AdvertisementMysql{db: db}
}
