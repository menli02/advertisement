package server

import (
	"context"

	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/logic"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/svc"
)

type AuthServer struct {
	svcCtx *svc.ServiceContext
}

func NewAuthServer(svcCtx *svc.ServiceContext) *AuthServer {
	return &AuthServer{svcCtx: svcCtx}
}

func (s *AuthServer) SignIn(ctx context.Context, in *auth.SignInRequest) (*auth.SignInResponse, error) {
	l := logic.NewSignInLogic(ctx, s.svcCtx)
	return l.SignIn(in)
}

func (s *AuthServer) RenewToken(ctx context.Context, in *auth.RenewTokenRequest) (*auth.RenewTokenResponse, error) {
	l := logic.NewRenewTokenLogic(ctx, s.svcCtx)
	return l.RenewToken(in)
}
