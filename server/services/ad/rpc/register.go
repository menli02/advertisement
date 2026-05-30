package rpc

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"

	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
)

type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error)      { return json.Marshal(v) }
func (jsonCodec) Unmarshal(data []byte, v interface{}) error { return json.Unmarshal(data, v) }
func (jsonCodec) Name() string                               { return "proto" }

func init() { encoding.RegisterCodec(jsonCodec{}) }

func RegisterAdServiceServer(s *grpc.Server, srv ad.AdServiceServer) {
	desc := &grpc.ServiceDesc{
		ServiceName: "ad.AdService",
		HandlerType: (*ad.AdServiceServer)(nil),
		Methods:     []grpc.MethodDesc{},
		Streams:     []grpc.StreamDesc{},
	}

	methods := []struct {
		name    string
		handler func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error)
	}{
		{
			"CreateAd",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.CreateAdRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).CreateAd(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/CreateAd"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).CreateAd(ctx, req.(*ad.CreateAdRequest))
				})
			},
		},
		{
			"GetAd",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.GetAdRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).GetAd(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/GetAd"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).GetAd(ctx, req.(*ad.GetAdRequest))
				})
			},
		},
		{
			"ListAds",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.ListAdsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).ListAds(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/ListAds"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).ListAds(ctx, req.(*ad.ListAdsRequest))
				})
			},
		},
		{
			"UpdateAd",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.UpdateAdRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).UpdateAd(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/UpdateAd"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).UpdateAd(ctx, req.(*ad.UpdateAdRequest))
				})
			},
		},
		{
			"DeleteAd",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.DeleteAdRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).DeleteAd(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/DeleteAd"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).DeleteAd(ctx, req.(*ad.DeleteAdRequest))
				})
			},
		},
		{
			"IncrementView",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.IncrementViewRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).IncrementView(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/IncrementView"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).IncrementView(ctx, req.(*ad.IncrementViewRequest))
				})
			},
		},
		{
			"CreateCategory",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.CreateCategoryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).CreateCategory(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/CreateCategory"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).CreateCategory(ctx, req.(*ad.CreateCategoryRequest))
				})
			},
		},
		{
			"ListCategories",
			func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ad.ListCategoriesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(ad.AdServiceServer).ListCategories(ctx, in)
				}
				info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/ad.AdService/ListCategories"}
				return interceptor(ctx, in, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(ad.AdServiceServer).ListCategories(ctx, req.(*ad.ListCategoriesRequest))
				})
			},
		},
	}

	for _, m := range methods {
		m := m
		desc.Methods = append(desc.Methods, grpc.MethodDesc{
			MethodName: m.name,
			Handler:    m.handler,
		})
	}

	s.RegisterService(desc, srv)
}
