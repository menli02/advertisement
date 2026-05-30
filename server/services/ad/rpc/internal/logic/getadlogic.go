package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type GetAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdLogic {
	return &GetAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *GetAdLogic) GetAd(in *ad.GetAdRequest) (*ad.AdResponse, error) {
	if in.Id == 0 && in.Slug == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	// Try cache first (by ID)
	if in.Id > 0 {
		cached, err := l.svcCtx.Repo.AdCache.Get(l.ctx, in.Id)
		if err == nil && cached != nil {
			return cached, nil
		}
		if err != nil && err != redis.Nil {
			l.Errorf("GetAd: cache get: %v", err)
		}
	}

	var result *ad.AdResponse

	if in.Id > 0 {
		a, err := l.svcCtx.Repo.Ad.GetByID(l.ctx, in.Id)
		if err != nil {
			l.Errorf("GetAd: db get by id: %v", err)
			return nil, errorcode.ErrInternal
		}
		if a == nil {
			return nil, errorcode.ErrAdNotFound
		}
		imgs, _ := l.svcCtx.Repo.Image.GetByAdID(l.ctx, a.ID)
		urls := make([]string, len(imgs))
		for i, img := range imgs {
			urls[i] = img.URL
		}
		result = adToResponse(a, urls)
	} else {
		a, err := l.svcCtx.Repo.Ad.GetBySlug(l.ctx, in.Slug)
		if err != nil {
			l.Errorf("GetAd: db get by slug: %v", err)
			return nil, errorcode.ErrInternal
		}
		if a == nil {
			return nil, errorcode.ErrAdNotFound
		}
		imgs, _ := l.svcCtx.Repo.Image.GetByAdID(l.ctx, a.ID)
		urls := make([]string, len(imgs))
		for i, img := range imgs {
			urls[i] = img.URL
		}
		result = adToResponse(a, urls)
	}

	// Populate cache
	if err := l.svcCtx.Repo.AdCache.Set(l.ctx, result); err != nil {
		l.Errorf("GetAd: cache set: %v", err)
	}

	return result, nil
}
