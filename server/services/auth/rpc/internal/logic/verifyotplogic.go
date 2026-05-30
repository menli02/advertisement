package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	commonjwt "github.com/menli02/advertisement/server/common/jwt"
	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/constant"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/svc"
)

type VerifyOTPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyOTPLogic {
	return &VerifyOTPLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyOTPLogic) VerifyOTP(in *auth.VerifyOTPRequest) (*auth.VerifyOTPResponse, error) {
	if in.OtpId == "" || in.Code == 0 || in.Phone == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	data, err := l.svcCtx.Repo.OTP.Get(l.ctx, in.OtpId)
	if err != nil {
		if err == redis.Nil {
			return nil, errorcode.ErrOTPExpired
		}
		l.Errorf("VerifyOTP: cache get: %v", err)
		return nil, errorcode.ErrInternal
	}

	if data.Phone != in.Phone {
		return nil, errorcode.ErrOTPInvalid
	}

	if data.Tries >= constant.OTPMaxTries {
		_ = l.svcCtx.Repo.OTP.Delete(l.ctx, in.OtpId)
		return nil, errorcode.ErrOTPMaxTries
	}

	if data.Code != in.Code {
		if err := l.svcCtx.Repo.OTP.IncrTries(l.ctx, in.OtpId, data); err != nil {
			l.Errorf("VerifyOTP: incr tries: %v", err)
		}
		return nil, errorcode.ErrOTPInvalid
	}

	_ = l.svcCtx.Repo.OTP.Delete(l.ctx, in.OtpId)

	now := time.Now()
	claims := &commonjwt.Claims{
		UserId:    0,
		UserType:  "user",
		TokenType: "temp",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   in.Phone,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(commonjwt.TempTokenExpiry)),
		},
	}
	tempToken, err := commonjwt.GenerateToken(claims, l.svcCtx.Config.App.TempJWTSecret)
	if err != nil {
		l.Errorf("VerifyOTP: generate temp token: %v", err)
		return nil, errorcode.ErrInternal
	}

	return &auth.VerifyOTPResponse{TempToken: tempToken}, nil
}
