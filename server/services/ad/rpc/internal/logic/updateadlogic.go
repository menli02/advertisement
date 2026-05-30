package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/models"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type UpdateAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdLogic {
	return &UpdateAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *UpdateAdLogic) UpdateAd(in *ad.UpdateAdRequest) (*ad.AdResponse, error) {
	if in.Id == 0 || in.UserId == 0 {
		return nil, errorcode.ErrInvalidArgument
	}

	existing, err := l.svcCtx.Repo.Ad.GetByID(l.ctx, in.Id)
	if err != nil {
		return nil, errorcode.ErrInternal
	}
	if existing == nil {
		return nil, errorcode.ErrAdNotFound
	}
	if existing.UserID != in.UserId {
		return nil, errorcode.ErrAdForbidden
	}

	status := existing.Status
	if in.Status != "" {
		status = in.Status
	}
	currency := existing.Currency
	if in.Currency != "" {
		currency = in.Currency
	}
	categoryID := existing.CategoryID
	if in.CategoryId > 0 {
		categoryID = in.CategoryId
	}

	updated, err := l.svcCtx.Repo.Ad.Update(l.ctx, &models.Advertisement{
		ID:          in.Id,
		UserID:      in.UserId,
		CategoryID:  categoryID,
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		Currency:    currency,
		Status:      status,
	})
	if err != nil {
		l.Errorf("UpdateAd: db update: %v", err)
		return nil, errorcode.ErrInternal
	}
	if updated == nil {
		return nil, errorcode.ErrAdNotFound
	}

	if len(in.Images) > 0 {
		if err := l.svcCtx.Repo.Image.ReplaceAll(l.ctx, updated.ID, in.Images); err != nil {
			l.Errorf("UpdateAd: replace images: %v", err)
		}
	}

	// Invalidate cache
	_ = l.svcCtx.Repo.AdCache.Delete(l.ctx, updated.ID)

	imgs, _ := l.svcCtx.Repo.Image.GetByAdID(l.ctx, updated.ID)
	urls := make([]string, len(imgs))
	for i, img := range imgs {
		urls[i] = img.URL
	}

	return adToResponse(updated, urls), nil
}
