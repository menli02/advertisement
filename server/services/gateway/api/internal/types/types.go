package types

// Auth requests/responses

type SendOTPRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type SendOTPResponse struct {
	OtpId  string `json:"otp_id"`
	ExpSec int32  `json:"exp_sec"`
}

type VerifyOTPRequest struct {
	OtpId string `json:"otp_id" validate:"required"`
	Code  int32  `json:"code" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

type VerifyOTPResponse struct {
	TempToken string `json:"temp_token"`
}

type SignInRequest struct {
	TempToken string `json:"temp_token" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserId       int64  `json:"user_id"`
}

type RenewTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RenewTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Ad requests/responses

type CreateAdRequest struct {
	CategoryId  int64    `json:"category_id" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Currency    string   `json:"currency"`
	Images      []string `json:"images"`
}

type UpdateAdRequest struct {
	CategoryId  int64    `json:"category_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Currency    string   `json:"currency"`
	Status      string   `json:"status"`
	Images      []string `json:"images"`
}

type ListAdsQuery struct {
	CategoryId int64  `form:"category_id,optional"`
	UserId     int64  `form:"user_id,optional"`
	Query      string `form:"query,optional"`
	SortBy     string `form:"sort_by,optional"`
	SortOrder  string `form:"sort_order,optional"`
	Page       int32  `form:"page,optional"`
	PageSize   int32  `form:"page_size,optional"`
}

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

type ListAdsResponse struct {
	Ads   []AdResponse `json:"ads"`
	Total int64        `json:"total"`
}

type CategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type ListCategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}
