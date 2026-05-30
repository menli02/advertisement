package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	authrpc "github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *SignInLogic) SignIn(req *types.SignInRequest) (*types.SignInResponse, error) {
	resp, err := l.svcCtx.Auth.SignIn(l.ctx, &authrpc.SignInRequest{
		TempToken: req.TempToken,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, err
	}
	return &types.SignInResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		UserId:       resp.UserId,
	}, nil
}
