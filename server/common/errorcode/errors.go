package errorcode

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidArgument = status.Error(400, "INVALID_ARGUMENT")
	ErrRequestInvalid  = status.Error(400, "REQUEST_INVALID")
	ErrUnauthorized    = status.Error(401, "UNAUTHORIZED")
	ErrTokenInvalid    = status.Error(401, "TOKEN_INVALID")
	ErrForbidden       = status.Error(403, "FORBIDDEN")
	ErrNotFound        = status.Error(404, "NOT_FOUND")
	ErrTooManyRequests = status.Error(429, "TOO_MANY_REQUESTS")
	ErrInternal        = status.Error(500, "INTERNAL_SERVER_ERROR")
	ErrInternalServer  = ErrInternal

	// Auth
	ErrPhoneNumberInvalid   = status.Error(400, "PHONE_NUMBER_INVALID")
	ErrOTPInvalid           = status.Error(401, "OTP_INVALID")
	ErrOTPExpired           = status.Error(401, "OTP_EXPIRED")
	ErrOTPMaxTries          = status.Error(429, "OTP_MAX_TRIES_EXCEEDED")
	ErrOTPRequestIdNotFound = status.Error(404, "OTP_REQUEST_ID_NOT_FOUND")

	// Advertisement
	ErrAdNotFound       = status.Error(404, "ADVERTISEMENT_NOT_FOUND")
	ErrAdForbidden      = status.Error(403, "ADVERTISEMENT_ACCESS_FORBIDDEN")
	ErrCategoryNotFound = status.Error(404, "CATEGORY_NOT_FOUND")
	ErrSlugTaken        = status.Error(409, "SLUG_ALREADY_TAKEN")
)

// HandleError maps gRPC status codes to HTTP responses.
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	st, ok := status.FromError(err)
	if !ok {
		httpx.Error(w, err)
		return
	}

	code := int(st.Code())
	switch code {
	case 400:
		httpx.WriteJson(w, http.StatusBadRequest, map[string]string{"code": st.Message()})
	case 401:
		httpx.WriteJson(w, http.StatusUnauthorized, map[string]string{"code": st.Message()})
	case 403:
		httpx.WriteJson(w, http.StatusForbidden, map[string]string{"code": st.Message()})
	case 404:
		httpx.WriteJson(w, http.StatusNotFound, map[string]string{"code": st.Message()})
	case 409:
		httpx.WriteJson(w, http.StatusConflict, map[string]string{"code": st.Message()})
	case 429:
		httpx.WriteJson(w, http.StatusTooManyRequests, map[string]string{"code": st.Message()})
	default:
		httpx.WriteJson(w, http.StatusInternalServerError, map[string]string{"code": st.Message()})
	}
}
