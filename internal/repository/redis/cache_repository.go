package redis

import (
	"Advertisement/internal/domain"
	"context"
	"github.com/go-redis/redis/v8"
)

type AdvertisementCache interface {
	CreateAdvertisement(ctx context.Context, ad *domain.Advertisement) error
	GetAdvertisementByID(ctx context.Context, id int) (*domain.Advertisement, error)
	GetAllSortedAndPaginated(ctx context.Context, sortBy, sortOrder string, offset, limit int) ([]*domain.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, id int, ad *domain.UpdateAdvertisementInput) error
	DeleteAdvertisement(ctx context.Context, id int) error
}
type AdvertisementRedis struct {
	client *redis.Client
}

func NewAdvertisementRedis(client *redis.Client) *AdvertisementRedis {
	return &AdvertisementRedis{
		client: client,
	}
}
