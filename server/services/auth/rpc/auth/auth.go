// Manual interfaces — replace with protoc-generated code after `make generate-auth`.
package auth

import (
	"context"

	"google.golang.org/grpc"
)

type SendOTPRequest  struct{ Phone string }
type SendOTPResponse struct{ OtpId string; ExpSec int32 }

type VerifyOTPRequest  struct{ OtpId string; Code int32; Phone string }
type VerifyOTPResponse struct{ TempToken string }

type SignInRequest  struct{ TempToken string; FirstName string; LastName string }
type SignInResponse struct{ AccessToken string; RefreshToken string; UserId int64 }

type RenewTokenRequest  struct{ RefreshToken string }
type RenewTokenResponse struct{ AccessToken string; RefreshToken string }

type OTPManagerClient interface {
	SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*SendOTPResponse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
}
type OTPManagerServer interface {
	SendOTP(context.Context, *SendOTPRequest) (*SendOTPResponse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
}

type AuthClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	RenewToken(ctx context.Context, in *RenewTokenRequest, opts ...grpc.CallOption) (*RenewTokenResponse, error)
}
type AuthServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	RenewToken(context.Context, *RenewTokenRequest) (*RenewTokenResponse, error)
}
