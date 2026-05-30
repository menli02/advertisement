package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/clients"
	"github.com/menli02/advertisement/server/services/gateway/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Auth   *clients.AuthClient
	Ad     *clients.AdClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   clients.NewAuthClient(zrpc.MustNewClient(c.AuthRpc)),
		Ad:     clients.NewAdClient(zrpc.MustNewClient(c.AdRpc)),
	}
}
