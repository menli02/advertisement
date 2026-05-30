package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	authrpc "github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type RenewTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRenewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenewTokenLogic {
	return &RenewTokenLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *RenewTokenLogic) RenewToken(req *types.RenewTokenRequest) (*types.RenewTokenResponse, error) {
	resp, err := l.svcCtx.Auth.RenewToken(l.ctx, &authrpc.RenewTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}
	return &types.RenewTokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}
