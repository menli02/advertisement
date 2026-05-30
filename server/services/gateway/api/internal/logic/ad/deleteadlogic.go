package ad

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	commonjwt "github.com/menli02/advertisement/server/common/jwt"
	"github.com/menli02/advertisement/server/common/errorcode"
	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
)

type DeleteAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdLogic {
	return &DeleteAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *DeleteAdLogic) DeleteAd(id int64) error {
	claims, ok := commonjwt.ClaimsFromContext(l.ctx)
	if !ok {
		return errorcode.ErrUnauthorized
	}

	_, err := l.svcCtx.Ad.DeleteAd(l.ctx, &adrpc.DeleteAdRequest{
		Id:     id,
		UserId: claims.UserId,
	})
	return err
}
