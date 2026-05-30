package logic

import (
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/models"
)

func adToResponse(a *models.Advertisement, images []string) *ad.AdResponse {
	return &ad.AdResponse{
		Id:          a.ID,
		UserId:      a.UserID,
		CategoryId:  a.CategoryID,
		Title:       a.Title,
		Description: a.Description,
		Slug:        a.Slug,
		Price:       a.Price,
		Currency:    a.Currency,
		Status:      a.Status,
		ViewCount:   a.ViewCount,
		Images:      images,
		CreatedAt:   a.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   a.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func categoryToResponse(c *models.Category) *ad.CategoryResponse {
	return &ad.CategoryResponse{
		Id:   c.ID,
		Name: c.Name,
		Slug: c.Slug,
	}
}
