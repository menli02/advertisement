package ad

import (
	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

func adResponseToType(r *adrpc.AdResponse) *types.AdResponse {
	if r == nil {
		return nil
	}
	return &types.AdResponse{
		Id:          r.Id,
		UserId:      r.UserId,
		CategoryId:  r.CategoryId,
		Title:       r.Title,
		Description: r.Description,
		Slug:        r.Slug,
		Price:       r.Price,
		Currency:    r.Currency,
		Status:      r.Status,
		ViewCount:   r.ViewCount,
		Images:      r.Images,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
