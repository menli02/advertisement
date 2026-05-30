package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	rpc "github.com/menli02/advertisement/server/services/auth/rpc"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/config"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/server"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/svc"
)

var configFile = flag.String("f", "configs/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterOTPManagerServer(grpcServer, server.NewOTPManagerServer(ctx))
		rpc.RegisterAuthServer(grpcServer, server.NewAuthServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting auth rpc server at %s...\n", c.ListenOn)
	s.Start()
}
