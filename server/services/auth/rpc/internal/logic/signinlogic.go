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

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *auth.SignInRequest) (*auth.SignInResponse, error) {
	if in.TempToken == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	tempClaims, err := commonjwt.ParseTempToken(in.TempToken, l.svcCtx.Config.App.TempJWTSecret)
	if err != nil {
		return nil, errorcode.ErrUnauthorized
	}

	phone := tempClaims.Subject

	user, err := l.svcCtx.Repo.User.Upsert(l.ctx, phone, in.FirstName, in.LastName)
	if err != nil {
		l.Errorf("SignIn: upsert user: %v", err)
		return nil, errorcode.ErrInternal
	}

	now := time.Now()

	accessClaims := &commonjwt.Claims{
		UserId:    user.ID,
		UserType:  "user",
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   phone,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(commonjwt.AccessTokenExpiry)),
		},
	}
	accessToken, err := commonjwt.GenerateToken(accessClaims, l.svcCtx.Config.App.JWTSecret)
	if err != nil {
		l.Errorf("SignIn: generate access token: %v", err)
		return nil, errorcode.ErrInternal
	}

	refreshClaims := &commonjwt.Claims{
		UserId:    user.ID,
		UserType:  "user",
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   phone,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(commonjwt.RefreshTokenExpiry)),
		},
	}
	refreshToken, err := commonjwt.GenerateToken(refreshClaims, l.svcCtx.Config.App.RefreshJWTSecret)
	if err != nil {
		l.Errorf("SignIn: generate refresh token: %v", err)
		return nil, errorcode.ErrInternal
	}

	return &auth.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       user.ID,
	}, nil
}
