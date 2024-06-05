package service

import (
	"Advertisement/internal/domain"
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"log"
)

var tracer = otel.GetTracerProvider().Tracer("advertisement")

func (s *Service) CreateAdvertisement(ctx context.Context, ad *domain.Advertisement) error {
	ctx, span := tracer.Start(ctx, "Service.CreateAdvertisement")
	defer span.End()
	err := s.repo.CreateAdvertisement(ctx, ad)
	if err != nil {
		return err
	}

	err = s.cache.CreateAdvertisement(ctx, ad)
	if err != nil {

		log.Printf("failed to update cache: %v", err)
	}

	return nil
}

func (s *Service) GetAdvertisementByID(ctx context.Context, id int) (*domain.Advertisement, error) {
	ctx, span := tracer.Start(ctx, "Service.GetAdvertisementByID")
	defer span.End()
	ad, err := s.cache.GetAdvertisementByID(ctx, id)
	if err != nil {
		ad, err := s.repo.GetAdvertisementByID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("advertisement is not found or cannot got from in database: %v", err)
		}
		err = s.cache.CreateAdvertisement(ctx, ad)
		if err != nil {
			return nil, fmt.Errorf("error while fetching advertisement from db to redis: %s", err)
		}
		return ad, nil
	}
	return ad, nil
}
func (s *Service) GetAllSortedAndPaginated(ctx context.Context, sortBy, sortOrder string, offset, limit int) ([]*domain.Advertisement, error) {
	ctx, span := tracer.Start(ctx, "Service.GetAllAdvertisement")
	defer span.End()
	ads, err := s.cache.GetAllSortedAndPaginated(ctx, sortBy, sortOrder, offset, limit)
	if err != nil {
		ads, err := s.repo.GetAllSortedAndPaginated(ctx, sortBy, sortOrder, offset, limit)
		if err != nil {
			return nil, fmt.Errorf("advertisements are not found in database : %s", err)
		}
		for _, ad := range ads {
			err = s.cache.CreateAdvertisement(ctx, ad)
			if err != nil {
				fmt.Printf("Error caching advertisement with ID %d: %s\n", ad.Id, err)
			}
		}
		return ads, nil
	}
	return ads, nil
}
func (s *Service) UpdateAdvertisement(ctx context.Context, id int, ad *domain.UpdateAdvertisementInput) error {
	ctx, span := tracer.Start(ctx, "Service.UpdateAdvertisement")
	defer span.End()
	if err := s.repo.UpdateAdvertisement(ctx, id, ad); err != nil {
		return fmt.Errorf("failed to update advertisement in MySQL: %v", err)
	}

	if err := s.cache.UpdateAdvertisement(ctx, id, ad); err != nil {
		return fmt.Errorf("failed to update advertisement in Redis: %v", err)
	}

	return nil
}

func (s *Service) DeleteAdvertisement(ctx context.Context, id int) error {
	ctx, span := tracer.Start(ctx, "DeleteAdvertisement")
	defer span.End()
	if err := s.repo.DeleteAdvertisement(ctx, id); err != nil {
		return fmt.Errorf("failed to delete advertisement from MySQL: %v", err)
	}
	if err := s.cache.DeleteAdvertisement(ctx, id); err != nil {
		return fmt.Errorf("failed to delete advertisement from REDIS: %v", err)
	}
	return nil
}
