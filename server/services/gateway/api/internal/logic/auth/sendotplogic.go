package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type SendOTPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendOTPLogic {
	return &SendOTPLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *SendOTPLogic) SendOTP(req *types.SendOTPRequest) (*types.SendOTPResponse, error) {
	resp, err := l.svcCtx.Auth.SendOTP(l.ctx, &auth.SendOTPRequest{Phone: req.Phone})
	if err != nil {
		return nil, err
	}
	return &types.SendOTPResponse{OtpId: resp.OtpId, ExpSec: resp.ExpSec}, nil
}
