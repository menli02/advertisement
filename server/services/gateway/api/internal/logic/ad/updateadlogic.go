package ad

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	commonjwt "github.com/menli02/advertisement/server/common/jwt"
	"github.com/menli02/advertisement/server/common/errorcode"
	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type UpdateAdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdLogic {
	return &UpdateAdLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *UpdateAdLogic) UpdateAd(id int64, req *types.UpdateAdRequest) (*types.AdResponse, error) {
	claims, ok := commonjwt.ClaimsFromContext(l.ctx)
	if !ok {
		return nil, errorcode.ErrUnauthorized
	}

	resp, err := l.svcCtx.Ad.UpdateAd(l.ctx, &adrpc.UpdateAdRequest{
		Id:          id,
		UserId:      claims.UserId,
		CategoryId:  req.CategoryId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Currency:    req.Currency,
		Status:      req.Status,
		Images:      req.Images,
	})
	if err != nil {
		return nil, err
	}

	return adResponseToType(resp), nil
}
