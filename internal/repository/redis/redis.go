package redis

import (
	"Advertisement/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"strconv"
	"time"
)

var tracer = otel.GetTracerProvider().Tracer("advertisement")

func (r *AdvertisementRedis) CreateAdvertisement(ctx context.Context, ad *domain.Advertisement) error {
	ctx, span := tracer.Start(ctx, "Create-redis")
	defer span.End()
	data, err := json.Marshal(ad)
	if err != nil {
		return fmt.Errorf("failed to serialize advertisement: %v", err)
	}

	err = r.client.Set(ctx, fmt.Sprintf("advertisement:%d", ad.Id), data, time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to add advertisement to Redis: %v", err)
	}

	return nil
}
func (r *AdvertisementRedis) GetAdvertisementByID(ctx context.Context, id int) (*domain.Advertisement, error) {
	ctx, span := tracer.Start(ctx, "GetById-redis")
	defer span.End()
	key := fmt.Sprintf("advertisement:%d", id)

	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var ad domain.Advertisement
	if err := json.Unmarshal([]byte(val), &ad); err != nil {
		return nil, fmt.Errorf("failed to unmarshal advertisement: %w", err)
	}

	return &ad, nil
}
func (r *AdvertisementRedis) GetAllSortedAndPaginated(ctx context.Context, sortBy, sortOrder string, offset, limit int) ([]*domain.Advertisement, error) {
	ctx, span := tracer.Start(ctx, "GetAll-redis")
	defer span.End()
	key := "advertisements:" + sortBy + ":" + sortOrder + ":offset:" + strconv.Itoa(offset) + ":limit:" + strconv.Itoa(limit)
	adsJSON, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var ads []*domain.Advertisement
	if err := json.Unmarshal([]byte(adsJSON), &ads); err != nil {
		return nil, err
	}

	return ads, nil
}
func (r *AdvertisementRedis) UpdateAdvertisement(ctx context.Context, id int, ad *domain.UpdateAdvertisementInput) error {
	ctx, span := tracer.Start(ctx, "Update-redis")
	defer span.End()
	updatedAdJSON, err := json.Marshal(ad)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("advertisement:%d", id)
	_, err = r.client.Set(ctx, key, updatedAdJSON, time.Hour).Result()
	if err != nil {
		return err
	}

	return nil
}
func (r *AdvertisementRedis) DeleteAdvertisement(ctx context.Context, id int) error {
	ctx, span := tracer.Start(ctx, "Delete-redis")
	defer span.End()
	key := fmt.Sprintf("advertisement:%d", id)
	_, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
