package logic

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/constant"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/models"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/svc"
)

type SendOTPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendOTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendOTPLogic {
	return &SendOTPLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendOTPLogic) SendOTP(in *auth.SendOTPRequest) (*auth.SendOTPResponse, error) {
	if in.Phone == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	code := rand.Intn(900000) + 100000
	otpID := uuid.NewString()

	data := &models.OTPData{
		Code:  int32(code),
		Phone: in.Phone,
		Tries: 0,
	}
	if err := l.svcCtx.Repo.OTP.Set(l.ctx, otpID, data); err != nil {
		l.Errorf("SendOTP: cache set: %v", err)
		return nil, errorcode.ErrInternal
	}

	// In production: send SMS via provider. For now, log it.
	l.Infof("OTP for %s: %s", in.Phone, fmt.Sprintf("%06d", code))

	return &auth.SendOTPResponse{
		OtpId:  otpID,
		ExpSec: int32(constant.OTPExpiry),
	}, nil
}
