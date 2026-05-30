package rpc

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"

	"github.com/menli02/advertisement/server/services/auth/rpc/auth"
)

// jsonCodec shadows the default protobuf codec so plain Go structs work over gRPC.
type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error)      { return json.Marshal(v) }
func (jsonCodec) Unmarshal(data []byte, v interface{}) error { return json.Unmarshal(data, v) }
func (jsonCodec) Name() string                               { return "proto" }

func init() { encoding.RegisterCodec(jsonCodec{}) }

// RegisterOTPManagerServer registers the OTPManager service with the gRPC server.
func RegisterOTPManagerServer(s *grpc.Server, srv auth.OTPManagerServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "auth.OTPManager",
		HandlerType: (*auth.OTPManagerServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SendOTP",
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					in := new(auth.SendOTPRequest)
					if err := dec(in); err != nil {
						return nil, err
					}
					if interceptor == nil {
						return srv.(auth.OTPManagerServer).SendOTP(ctx, in)
					}
					info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.OTPManager/SendOTP"}
					handler := func(ctx context.Context, req interface{}) (interface{}, error) {
						return srv.(auth.OTPManagerServer).SendOTP(ctx, req.(*auth.SendOTPRequest))
					}
					return interceptor(ctx, in, info, handler)
				},
			},
			{
				MethodName: "VerifyOTP",
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					in := new(auth.VerifyOTPRequest)
					if err := dec(in); err != nil {
						return nil, err
					}
					if interceptor == nil {
						return srv.(auth.OTPManagerServer).VerifyOTP(ctx, in)
					}
					info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.OTPManager/VerifyOTP"}
					handler := func(ctx context.Context, req interface{}) (interface{}, error) {
						return srv.(auth.OTPManagerServer).VerifyOTP(ctx, req.(*auth.VerifyOTPRequest))
					}
					return interceptor(ctx, in, info, handler)
				},
			},
		},
		Streams: []grpc.StreamDesc{},
	}, srv)
}

// RegisterAuthServer registers the Auth service with the gRPC server.
func RegisterAuthServer(s *grpc.Server, srv auth.AuthServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: "auth.Auth",
		HandlerType: (*auth.AuthServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SignIn",
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					in := new(auth.SignInRequest)
					if err := dec(in); err != nil {
						return nil, err
					}
					if interceptor == nil {
						return srv.(auth.AuthServer).SignIn(ctx, in)
					}
					info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.Auth/SignIn"}
					handler := func(ctx context.Context, req interface{}) (interface{}, error) {
						return srv.(auth.AuthServer).SignIn(ctx, req.(*auth.SignInRequest))
					}
					return interceptor(ctx, in, info, handler)
				},
			},
			{
				MethodName: "RenewToken",
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					in := new(auth.RenewTokenRequest)
					if err := dec(in); err != nil {
						return nil, err
					}
					if interceptor == nil {
						return srv.(auth.AuthServer).RenewToken(ctx, in)
					}
					info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.Auth/RenewToken"}
					handler := func(ctx context.Context, req interface{}) (interface{}, error) {
						return srv.(auth.AuthServer).RenewToken(ctx, req.(*auth.RenewTokenRequest))
					}
					return interceptor(ctx, in, info, handler)
				},
			},
		},
		Streams: []grpc.StreamDesc{},
	}, srv)
}
