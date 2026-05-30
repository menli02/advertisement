package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/menli02/advertisement/server/pkg/pgsqlx"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/config"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/repository"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/repository/cache"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/repository/db"
)

type ServiceContext struct {
	Config config.Config
	Repo   *repository.Repository
}

func NewServiceContext(c config.Config) *ServiceContext {
	pg := pgsqlx.NewPostgresSQL(pgsqlx.Config{DSN: c.Postgres.DSN})

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	return &ServiceContext{
		Config: c,
		Repo: &repository.Repository{
			User: db.NewUserRepository(pg.Pool()),
			OTP:  cache.NewOTPCache(rdb),
		},
	}
}
