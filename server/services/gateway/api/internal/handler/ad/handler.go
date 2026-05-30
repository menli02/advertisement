package ad

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/menli02/advertisement/server/common/errorcode"
	adlogic "github.com/menli02/advertisement/server/services/gateway/api/internal/logic/ad"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

type idPathVar struct {
	Id int64 `path:"id"`
}

func CreateAdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateAdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := adlogic.NewCreateAdLogic(r.Context(), svcCtx).CreateAd(&req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func GetAdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p idPathVar
		if err := httpx.Parse(r, &p); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		slug := r.URL.Query().Get("slug")

		resp, err := adlogic.NewGetAdLogic(r.Context(), svcCtx).GetAd(p.Id, slug)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func ListAdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var q types.ListAdsQuery
		if err := httpx.Parse(r, &q); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := adlogic.NewListAdsLogic(r.Context(), svcCtx).ListAds(&q)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func UpdateAdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p idPathVar
		if err := httpx.Parse(r, &p); err != nil || p.Id == 0 {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		var req types.UpdateAdRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := adlogic.NewUpdateAdLogic(r.Context(), svcCtx).UpdateAd(p.Id, &req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func DeleteAdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p idPathVar
		if err := httpx.Parse(r, &p); err != nil || p.Id == 0 {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		if err := adlogic.NewDeleteAdLogic(r.Context(), svcCtx).DeleteAd(p.Id); err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, map[string]bool{"success": true})
	}
}

func CreateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCategoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := adlogic.NewCategoryLogic(r.Context(), svcCtx).CreateCategory(&req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func ListCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := adlogic.NewCategoryLogic(r.Context(), svcCtx).ListCategories()
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}
