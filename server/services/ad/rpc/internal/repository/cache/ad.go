package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/menli02/advertisement/server/services/ad/rpc/ad"
)

const adCacheTTL = 10 * time.Minute

type AdCache struct {
	client *redis.Client
}

func NewAdCache(client *redis.Client) *AdCache {
	return &AdCache{client: client}
}

func adKey(id int64) string       { return fmt.Sprintf("ad:%d", id) }
func viewKey(id int64) string     { return fmt.Sprintf("ad:view:%d", id) }

func (c *AdCache) Set(ctx context.Context, a *ad.AdResponse) error {
	b, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return c.client.SetEx(ctx, adKey(a.Id), b, adCacheTTL).Err()
}

func (c *AdCache) Get(ctx context.Context, id int64) (*ad.AdResponse, error) {
	b, err := c.client.Get(ctx, adKey(id)).Bytes()
	if err != nil {
		return nil, err
	}
	var a ad.AdResponse
	if err := json.Unmarshal(b, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

func (c *AdCache) Delete(ctx context.Context, id int64) error {
	return c.client.Del(ctx, adKey(id)).Err()
}

// IncrView increments Redis view counter and returns the new value.
// Actual DB sync happens periodically or on cache eviction.
func (c *AdCache) IncrView(ctx context.Context, id int64) (int64, error) {
	pipe := c.client.Pipeline()
	incr := pipe.Incr(ctx, viewKey(id))
	pipe.Expire(ctx, viewKey(id), 24*time.Hour)
	if _, err := pipe.Exec(ctx); err != nil {
		return 0, err
	}
	return incr.Val(), nil
}

func (c *AdCache) GetViewCount(ctx context.Context, id int64) (int64, error) {
	val, err := c.client.Get(ctx, viewKey(id)).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return val, err
}
