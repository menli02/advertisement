package ad

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type ListAdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAdsLogic {
	return &ListAdsLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *ListAdsLogic) ListAds(q *types.ListAdsQuery) (*types.ListAdsResponse, error) {
	resp, err := l.svcCtx.Ad.ListAds(l.ctx, &adrpc.ListAdsRequest{
		CategoryId: q.CategoryId,
		UserId:     q.UserId,
		Query:      q.Query,
		SortBy:     q.SortBy,
		SortOrder:  q.SortOrder,
		Page:       q.Page,
		PageSize:   q.PageSize,
	})
	if err != nil {
		return nil, err
	}

	ads := make([]types.AdResponse, 0, len(resp.Ads))
	for _, a := range resp.Ads {
		if r := adResponseToType(a); r != nil {
			ads = append(ads, *r)
		}
	}

	return &types.ListAdsResponse{Ads: ads, Total: resp.Total}, nil
}
