package server

import (
	"context"

	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/logic"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/svc"
)

type OTPManagerServer struct {
	svcCtx *svc.ServiceContext
}

func NewOTPManagerServer(svcCtx *svc.ServiceContext) *OTPManagerServer {
	return &OTPManagerServer{svcCtx: svcCtx}
}

func (s *OTPManagerServer) SendOTP(ctx context.Context, in *auth.SendOTPRequest) (*auth.SendOTPResponse, error) {
	l := logic.NewSendOTPLogic(ctx, s.svcCtx)
	return l.SendOTP(in)
}

func (s *OTPManagerServer) VerifyOTP(ctx context.Context, in *auth.VerifyOTPRequest) (*auth.VerifyOTPResponse, error) {
	l := logic.NewVerifyOTPLogic(ctx, s.svcCtx)
	return l.VerifyOTP(in)
}
