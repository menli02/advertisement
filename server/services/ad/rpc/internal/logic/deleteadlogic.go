package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type DeleteAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdLogic {
	return &DeleteAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *DeleteAdLogic) DeleteAd(in *ad.DeleteAdRequest) (*ad.DeleteAdResponse, error) {
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

	ok, err := l.svcCtx.Repo.Ad.SoftDelete(l.ctx, in.Id, in.UserId)
	if err != nil {
		l.Errorf("DeleteAd: soft delete: %v", err)
		return nil, errorcode.ErrInternal
	}

	_ = l.svcCtx.Repo.AdCache.Delete(l.ctx, in.Id)

	return &ad.DeleteAdResponse{Success: ok}, nil
}
