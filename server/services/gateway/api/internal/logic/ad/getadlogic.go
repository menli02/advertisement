package ad

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type GetAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdLogic {
	return &GetAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *GetAdLogic) GetAd(id int64, slug string) (*types.AdResponse, error) {
	resp, err := l.svcCtx.Ad.GetAd(l.ctx, &adrpc.GetAdRequest{Id: id, Slug: slug})
	if err != nil {
		return nil, err
	}

	// fire-and-forget view increment
	go func() {
		_, _ = l.svcCtx.Ad.IncrementView(context.Background(), &adrpc.IncrementViewRequest{Id: resp.Id})
	}()

	return adResponseToType(resp), nil
}
