package logic

import (
	"context"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/models"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type CreateAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAdLogic {
	return &CreateAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *CreateAdLogic) CreateAd(in *ad.CreateAdRequest) (*ad.AdResponse, error) {
	if in.UserId == 0 || in.CategoryId == 0 || in.Title == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	cat, err := l.svcCtx.Repo.Category.GetByID(l.ctx, in.CategoryId)
	if err != nil {
		l.Errorf("CreateAd: get category: %v", err)
		return nil, errorcode.ErrInternal
	}
	if cat == nil {
		return nil, errorcode.ErrCategoryNotFound
	}

	// Generate unique slug
	base := slug.Make(in.Title)
	adSlug := base
	for i := 1; ; i++ {
		exists, err := l.svcCtx.Repo.Ad.SlugExists(l.ctx, adSlug)
		if err != nil {
			return nil, errorcode.ErrInternal
		}
		if !exists {
			break
		}
		adSlug = fmt.Sprintf("%s-%d", base, i)
	}

	newAd := &models.Advertisement{
		UserID:      in.UserId,
		CategoryID:  in.CategoryId,
		Title:       in.Title,
		Description: in.Description,
		Slug:        adSlug,
		Price:       in.Price,
		Currency:    in.Currency,
	}

	created, err := l.svcCtx.Repo.Ad.Create(l.ctx, newAd)
	if err != nil {
		l.Errorf("CreateAd: db create: %v", err)
		return nil, errorcode.ErrInternal
	}

	if len(in.Images) > 0 {
		if err := l.svcCtx.Repo.Image.ReplaceAll(l.ctx, created.ID, in.Images); err != nil {
			l.Errorf("CreateAd: replace images: %v", err)
		}
	}

	return adToResponse(created, in.Images), nil
}
