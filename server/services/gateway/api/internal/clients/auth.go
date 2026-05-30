package clients

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	authrpc "github.com/menli02/advertisement/server/services/auth/rpc/auth"
)

type AuthClient struct {
	conn zrpc.Client
}

func NewAuthClient(conn zrpc.Client) *AuthClient {
	return &AuthClient{conn: conn}
}

func (c *AuthClient) SendOTP(ctx context.Context, in *authrpc.SendOTPRequest, opts ...grpc.CallOption) (*authrpc.SendOTPResponse, error) {
	out := new(authrpc.SendOTPResponse)
	err := c.conn.Conn().Invoke(ctx, "/auth.OTPManager/SendOTP", in, out, opts...)
	return out, err
}

func (c *AuthClient) VerifyOTP(ctx context.Context, in *authrpc.VerifyOTPRequest, opts ...grpc.CallOption) (*authrpc.VerifyOTPResponse, error) {
	out := new(authrpc.VerifyOTPResponse)
	err := c.conn.Conn().Invoke(ctx, "/auth.OTPManager/VerifyOTP", in, out, opts...)
	return out, err
}

func (c *AuthClient) SignIn(ctx context.Context, in *authrpc.SignInRequest, opts ...grpc.CallOption) (*authrpc.SignInResponse, error) {
	out := new(authrpc.SignInResponse)
	err := c.conn.Conn().Invoke(ctx, "/auth.Auth/SignIn", in, out, opts...)
	return out, err
}

func (c *AuthClient) RenewToken(ctx context.Context, in *authrpc.RenewTokenRequest, opts ...grpc.CallOption) (*authrpc.RenewTokenResponse, error) {
	out := new(authrpc.RenewTokenResponse)
	err := c.conn.Conn().Invoke(ctx, "/auth.Auth/RenewToken", in, out, opts...)
	return out, err
}
