package logic

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/menli02/advertisement/server/common/errorcode"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
	"github.com/menli02/advertisement/server/services/ad/rpc/internal/svc"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *CreateCategoryLogic) CreateCategory(in *ad.CreateCategoryRequest) (*ad.CategoryResponse, error) {
	if in.Name == "" {
		return nil, errorcode.ErrInvalidArgument
	}

	cat, err := l.svcCtx.Repo.Category.Create(l.ctx, in.Name, slug.Make(in.Name))
	if err != nil {
		l.Errorf("CreateCategory: db create: %v", err)
		return nil, errorcode.ErrInternal
	}

	return categoryToResponse(cat), nil
}
