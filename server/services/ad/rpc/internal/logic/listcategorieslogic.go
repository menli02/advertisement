package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type ListCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoriesLogic {
	return &ListCategoriesLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *ListCategoriesLogic) ListCategories(in *ad.ListCategoriesRequest) (*ad.ListCategoriesResponse, error) {
	cats, err := l.svcCtx.Repo.Category.List(l.ctx)
	if err != nil {
		l.Errorf("ListCategories: db list: %v", err)
		return nil, errorcode.ErrInternal
	}

	responses := make([]*ad.CategoryResponse, 0, len(cats))
	for _, c := range cats {
		responses = append(responses, categoryToResponse(c))
	}

	return &ad.ListCategoriesResponse{Categories: responses}, nil
}
