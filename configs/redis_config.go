package configs

import "github.com/go-redis/redis/v8"

type Options struct {
	Address  string
	Password string
	DB       int
}

func NewRedisCache(cfg Options) (*redis.Client, error) {
	redisOptions := &redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	rdb := redis.NewClient(redisOptions)
	return rdb, nil
}
