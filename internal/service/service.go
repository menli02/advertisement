package service

import (
	"Advertisement/internal/domain"
	"Advertisement/internal/repository/mysql"
	"Advertisement/internal/repository/redis"
	"context"
)

type Advertisement interface {
	CreateAdvertisement(ctx context.Context, ad *domain.Advertisement) error
	GetAdvertisementByID(ctx context.Context, id int) (*domain.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, id int, ad *domain.UpdateAdvertisementInput) error
	DeleteAdvertisement(ctx context.Context, id int) error
	GetAllSortedAndPaginated(ctx context.Context, sortBy, sortOrder string, offset, limit int) ([]*domain.Advertisement, error)
}

type Service struct {
	repo  mysql.Repository
	cache redis.AdvertisementCache
}

func NewService(redisCache redis.AdvertisementCache, mysqlRepo mysql.Repository) *Service {
	return &Service{
		cache: redisCache,
		repo:  mysqlRepo,
	}
}
