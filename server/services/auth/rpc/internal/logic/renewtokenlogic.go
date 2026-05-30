package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"

	commonjwt "github.com/menli02/advertisement/server/common/jwt"
	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/svc"
)

type RenewTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRenewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenewTokenLogic {
	return &RenewTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RenewTokenLogic) RenewToken(in *auth.RenewTokenRequest) (*auth.RenewTokenResponse, error) {
	if in.RefreshToken == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	claims, err := commonjwt.ParseToken(in.RefreshToken, l.svcCtx.Config.App.RefreshJWTSecret)
	if err != nil {
		return nil, errorcode.ErrUnauthorized
	}
	if claims.TokenType != "refresh" {
		return nil, errorcode.ErrUnauthorized
	}

	now := time.Now()

	accessClaims := &commonjwt.Claims{
		UserId:    claims.UserId,
		UserType:  claims.UserType,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   claims.Subject,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(commonjwt.AccessTokenExpiry)),
		},
	}
	accessToken, err := commonjwt.GenerateToken(accessClaims, l.svcCtx.Config.App.JWTSecret)
	if err != nil {
		l.Errorf("RenewToken: generate access token: %v", err)
		return nil, errorcode.ErrInternal
	}

	refreshClaims := &commonjwt.Claims{
		UserId:    claims.UserId,
		UserType:  claims.UserType,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   claims.Subject,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(commonjwt.RefreshTokenExpiry)),
		},
	}
	refreshToken, err := commonjwt.GenerateToken(refreshClaims, l.svcCtx.Config.App.RefreshJWTSecret)
	if err != nil {
		l.Errorf("RenewToken: generate refresh token: %v", err)
		return nil, errorcode.ErrInternal
	}

	return &auth.RenewTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
