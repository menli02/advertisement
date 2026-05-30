package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Postgres struct{ DSN string }
	Redis    struct{ Addr string; Password string; DB int }
	App      struct{ TempJWTSecret string; JWTSecret string; RefreshJWTSecret string }
}
