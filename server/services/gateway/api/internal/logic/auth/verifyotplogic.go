package auth

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	authrpc "github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type VerifyOTPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyOTPLogic {
	return &VerifyOTPLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *VerifyOTPLogic) VerifyOTP(req *types.VerifyOTPRequest) (*types.VerifyOTPResponse, error) {
	resp, err := l.svcCtx.Auth.VerifyOTP(l.ctx, &authrpc.VerifyOTPRequest{
		OtpId: req.OtpId,
		Code:  req.Code,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &types.VerifyOTPResponse{TempToken: resp.TempToken}, nil
}
