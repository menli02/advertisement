package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type ListAdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAdsLogic {
	return &ListAdsLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *ListAdsLogic) ListAds(in *ad.ListAdsRequest) (*ad.ListAdsResponse, error) {
	ads, total, err := l.svcCtx.Repo.Ad.List(
		l.ctx,
		in.CategoryId, in.UserId,
		in.Query, in.SortBy, in.SortOrder,
		in.Page, in.PageSize,
	)
	if err != nil {
		l.Errorf("ListAds: db list: %v", err)
		return nil, errorcode.ErrInternal
	}

	responses := make([]*ad.AdResponse, 0, len(ads))
	for _, a := range ads {
		imgs, _ := l.svcCtx.Repo.Image.GetByAdID(l.ctx, a.ID)
		urls := make([]string, len(imgs))
		for i, img := range imgs {
			urls[i] = img.URL
		}
		responses = append(responses, adToResponse(a, urls))
	}

	return &ad.ListAdsResponse{Ads: responses, Total: total}, nil
}
