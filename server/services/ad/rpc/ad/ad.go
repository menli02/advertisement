// Manual interfaces — replace with protoc-generated code after `make generate-ad`.
package ad

import (
	"context"

	"google.golang.org/grpc"
)

// --- Request/Response types ---

type AdResponse struct {
	Id          int64    `json:"id"`
	UserId      int64    `json:"user_id"`
	CategoryId  int64    `json:"category_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Price       float64  `json:"price"`
	Currency    string   `json:"currency"`
	Status      string   `json:"status"`
	ViewCount   int64    `json:"view_count"`
	Images      []string `json:"images"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type CreateAdRequest struct {
	UserId      int64    `json:"user_id"`
	CategoryId  int64    `json:"category_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Currency    string   `json:"currency"`
	Images      []string `json:"images"`
}

type GetAdRequest struct {
	Id   int64  `json:"id"`
	Slug string `json:"slug"`
}

type ListAdsRequest struct {
	CategoryId int64  `json:"category_id"`
	UserId     int64  `json:"user_id"`
	Query      string `json:"query"`
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order"`
	Page       int32  `json:"page"`
	PageSize   int32  `json:"page_size"`
}

type ListAdsResponse struct {
	Ads   []*AdResponse `json:"ads"`
	Total int64         `json:"total"`
}

type UpdateAdRequest struct {
	Id          int64    `json:"id"`
	UserId      int64    `json:"user_id"`
	CategoryId  int64    `json:"category_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Currency    string   `json:"currency"`
	Status      string   `json:"status"`
	Images      []string `json:"images"`
}

type DeleteAdRequest  struct{ Id int64 `json:"id"`; UserId int64 `json:"user_id"` }
type DeleteAdResponse struct{ Success bool `json:"success"` }

type IncrementViewRequest  struct{ Id int64 `json:"id"` }
type IncrementViewResponse struct{ ViewCount int64 `json:"view_count"` }

type CategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateCategoryRequest  struct{ Name string `json:"name"` }
type ListCategoriesRequest  struct{}
type ListCategoriesResponse struct{ Categories []*CategoryResponse `json:"categories"` }

// --- Server/Client interfaces ---

type AdServiceClient interface {
	CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*AdResponse, error)
	GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*AdResponse, error)
	ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (*ListAdsResponse, error)
	UpdateAd(ctx context.Context, in *UpdateAdRequest, opts ...grpc.CallOption) (*AdResponse, error)
	DeleteAd(ctx context.Context, in *DeleteAdRequest, opts ...grpc.CallOption) (*DeleteAdResponse, error)
	IncrementView(ctx context.Context, in *IncrementViewRequest, opts ...grpc.CallOption) (*IncrementViewResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
}

type AdServiceServer interface {
	CreateAd(context.Context, *CreateAdRequest) (*AdResponse, error)
	GetAd(context.Context, *GetAdRequest) (*AdResponse, error)
	ListAds(context.Context, *ListAdsRequest) (*ListAdsResponse, error)
	UpdateAd(context.Context, *UpdateAdRequest) (*AdResponse, error)
	DeleteAd(context.Context, *DeleteAdRequest) (*DeleteAdResponse, error)
	IncrementView(context.Context, *IncrementViewRequest) (*IncrementViewResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error)
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
}
