package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type IncrementViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrementViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrementViewLogic {
	return &IncrementViewLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *IncrementViewLogic) IncrementView(in *ad.IncrementViewRequest) (*ad.IncrementViewResponse, error) {
	if in.Id == 0 {
		return nil, errorcode.ErrInvalidArgument
	}

	// Increment in DB for durability, use Redis for fast reads
	count, err := l.svcCtx.Repo.Ad.IncrementView(l.ctx, in.Id)
	if err != nil {
		l.Errorf("IncrementView: db increment: %v", err)
		return nil, errorcode.ErrInternal
	}

	// Also sync Redis counter
	_, _ = l.svcCtx.Repo.AdCache.IncrView(l.ctx, in.Id)

	return &ad.IncrementViewResponse{ViewCount: count}, nil
}
