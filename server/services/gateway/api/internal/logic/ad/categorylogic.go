package ad

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	adrpc "github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type CategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *CategoryLogic) CreateCategory(req *types.CreateCategoryRequest) (*types.CategoryResponse, error) {
	resp, err := l.svcCtx.Ad.CreateCategory(l.ctx, &adrpc.CreateCategoryRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &types.CategoryResponse{Id: resp.Id, Name: resp.Name, Slug: resp.Slug}, nil
}

func (l *CategoryLogic) ListCategories() (*types.ListCategoriesResponse, error) {
	resp, err := l.svcCtx.Ad.ListCategories(l.ctx, &adrpc.ListCategoriesRequest{})
	if err != nil {
		return nil, err
	}
	cats := make([]types.CategoryResponse, 0, len(resp.Categories))
	for _, c := range resp.Categories {
		cats = append(cats, types.CategoryResponse{Id: c.Id, Name: c.Name, Slug: c.Slug})
	}
	return &types.ListCategoriesResponse{Categories: cats}, nil
}
