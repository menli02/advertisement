package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/menli02/advertisement/server/common/errorcode"
	authlogic "github.com/menli02/advertisement/server/services/gateway/api/internal/logic/auth"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/types"
)

func SendOTPHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendOTPRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := authlogic.NewSendOTPLogic(r.Context(), svcCtx).SendOTP(&req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func VerifyOTPHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyOTPRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := authlogic.NewVerifyOTPLogic(r.Context(), svcCtx).VerifyOTP(&req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func SignInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := authlogic.NewSignInLogic(r.Context(), svcCtx).SignIn(&req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}

func RenewTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RenewTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorcode.ErrInvalidArgument)
			return
		}
		resp, err := authlogic.NewRenewTokenLogic(r.Context(), svcCtx).RenewToken(&req)
		if err != nil {
			errorcode.HandleError(w, r, err)
			return
		}
		httpx.OkJson(w, resp)
	}
}
