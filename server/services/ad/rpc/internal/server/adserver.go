package server

import (
	"context"

	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/logic"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type AdServer struct {
	svcCtx *svc.ServiceContext
}

func NewAdServer(svcCtx *svc.ServiceContext) *AdServer {
	return &AdServer{svcCtx: svcCtx}
}

func (s *AdServer) CreateAd(ctx context.Context, in *ad.CreateAdRequest) (*ad.AdResponse, error) {
	return logic.NewCreateAdLogic(ctx, s.svcCtx).CreateAd(in)
}

func (s *AdServer) GetAd(ctx context.Context, in *ad.GetAdRequest) (*ad.AdResponse, error) {
	return logic.NewGetAdLogic(ctx, s.svcCtx).GetAd(in)
}

func (s *AdServer) ListAds(ctx context.Context, in *ad.ListAdsRequest) (*ad.ListAdsResponse, error) {
	return logic.NewListAdsLogic(ctx, s.svcCtx).ListAds(in)
}

func (s *AdServer) UpdateAd(ctx context.Context, in *ad.UpdateAdRequest) (*ad.AdResponse, error) {
	return logic.NewUpdateAdLogic(ctx, s.svcCtx).UpdateAd(in)
}

func (s *AdServer) DeleteAd(ctx context.Context, in *ad.DeleteAdRequest) (*ad.DeleteAdResponse, error) {
	return logic.NewDeleteAdLogic(ctx, s.svcCtx).DeleteAd(in)
}

func (s *AdServer) IncrementView(ctx context.Context, in *ad.IncrementViewRequest) (*ad.IncrementViewResponse, error) {
	return logic.NewIncrementViewLogic(ctx, s.svcCtx).IncrementView(in)
}

func (s *AdServer) CreateCategory(ctx context.Context, in *ad.CreateCategoryRequest) (*ad.CategoryResponse, error) {
	return logic.NewCreateCategoryLogic(ctx, s.svcCtx).CreateCategory(in)
}

func (s *AdServer) ListCategories(ctx context.Context, in *ad.ListCategoriesRequest) (*ad.ListCategoriesResponse, error) {
	return logic.NewListCategoriesLogic(ctx, s.svcCtx).ListCategories(in)
}
