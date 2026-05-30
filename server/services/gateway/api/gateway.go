package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	handlerAd   "github.com/menli02/advertisement/server/services/gateway/api/internal/handler/ad"
	handlerAuth "github.com/menli02/advertisement/server/services/gateway/api/internal/handler/auth"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/config"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/middleware"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/svc"
)

var configFile = flag.String("f", "configs/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	authMW := middleware.NewAuthMiddleware(c.App.JWTSecret)

	// Public auth routes
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/v1/auth/send-otp", Handler: handlerAuth.SendOTPHandler(ctx)})
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/v1/auth/verify-otp", Handler: handlerAuth.VerifyOTPHandler(ctx)})
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/v1/auth/sign-in", Handler: handlerAuth.SignInHandler(ctx)})
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/v1/auth/renew-token", Handler: handlerAuth.RenewTokenHandler(ctx)})

	// Public ad routes
	server.AddRoute(rest.Route{Method: http.MethodGet, Path: "/v1/ads", Handler: handlerAd.ListAdsHandler(ctx)})
	server.AddRoute(rest.Route{Method: http.MethodGet, Path: "/v1/ads/:id", Handler: handlerAd.GetAdHandler(ctx)})
	server.AddRoute(rest.Route{Method: http.MethodGet, Path: "/v1/categories", Handler: handlerAd.ListCategoriesHandler(ctx)})

	// Protected ad routes
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/v1/ads", Handler: authMW.Handle(handlerAd.CreateAdHandler(ctx))})
	server.AddRoute(rest.Route{Method: http.MethodPut, Path: "/v1/ads/:id", Handler: authMW.Handle(handlerAd.UpdateAdHandler(ctx))})
	server.AddRoute(rest.Route{Method: http.MethodDelete, Path: "/v1/ads/:id", Handler: authMW.Handle(handlerAd.DeleteAdHandler(ctx))})
	server.AddRoute(rest.Route{Method: http.MethodPost, Path: "/v1/categories", Handler: authMW.Handle(handlerAd.CreateCategoryHandler(ctx))})

	fmt.Printf("Starting gateway at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
