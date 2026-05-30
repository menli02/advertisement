package clients

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
)

type AdClient struct {
	conn zrpc.Client
}

func NewAdClient(conn zrpc.Client) *AdClient {
	return &AdClient{conn: conn}
}

func (c *AdClient) CreateAd(ctx context.Context, in *adrpc.CreateAdRequest, opts ...grpc.CallOption) (*adrpc.AdResponse, error) {
	out := new(adrpc.AdResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/CreateAd", in, out, opts...)
	return out, err
}

func (c *AdClient) GetAd(ctx context.Context, in *adrpc.GetAdRequest, opts ...grpc.CallOption) (*adrpc.AdResponse, error) {
	out := new(adrpc.AdResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/GetAd", in, out, opts...)
	return out, err
}

func (c *AdClient) ListAds(ctx context.Context, in *adrpc.ListAdsRequest, opts ...grpc.CallOption) (*adrpc.ListAdsResponse, error) {
	out := new(adrpc.ListAdsResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/ListAds", in, out, opts...)
	return out, err
}

func (c *AdClient) UpdateAd(ctx context.Context, in *adrpc.UpdateAdRequest, opts ...grpc.CallOption) (*adrpc.AdResponse, error) {
	out := new(adrpc.AdResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/UpdateAd", in, out, opts...)
	return out, err
}

func (c *AdClient) DeleteAd(ctx context.Context, in *adrpc.DeleteAdRequest, opts ...grpc.CallOption) (*adrpc.DeleteAdResponse, error) {
	out := new(adrpc.DeleteAdResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/DeleteAd", in, out, opts...)
	return out, err
}

func (c *AdClient) IncrementView(ctx context.Context, in *adrpc.IncrementViewRequest, opts ...grpc.CallOption) (*adrpc.IncrementViewResponse, error) {
	out := new(adrpc.IncrementViewResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/IncrementView", in, out, opts...)
	return out, err
}

func (c *AdClient) CreateCategory(ctx context.Context, in *adrpc.CreateCategoryRequest, opts ...grpc.CallOption) (*adrpc.CategoryResponse, error) {
	out := new(adrpc.CategoryResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/CreateCategory", in, out, opts...)
	return out, err
}

func (c *AdClient) ListCategories(ctx context.Context, in *adrpc.ListCategoriesRequest, opts ...grpc.CallOption) (*adrpc.ListCategoriesResponse, error) {
	out := new(adrpc.ListCategoriesResponse)
	err := c.conn.Conn().Invoke(ctx, "/ad.AdService/ListCategories", in, out, opts...)
	return out, err
}
