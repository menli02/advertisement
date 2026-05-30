package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/constant"
	"github.com/menli02/advertisement/server/services/auth/rpc/internal/models"
)

type OTPCache struct {
	client *redis.Client
}

func NewOTPCache(client *redis.Client) *OTPCache {
	return &OTPCache{client: client}
}

func otpKey(otpID string) string {
	return fmt.Sprintf("otp:%s", otpID)
}

func (c *OTPCache) Set(ctx context.Context, otpID string, data *models.OTPData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.SetEx(ctx, otpKey(otpID), b, time.Duration(constant.OTPExpiry)*time.Second).Err()
}

func (c *OTPCache) Get(ctx context.Context, otpID string) (*models.OTPData, error) {
	b, err := c.client.Get(ctx, otpKey(otpID)).Bytes()
	if err != nil {
		return nil, err
	}
	var data models.OTPData
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *OTPCache) IncrTries(ctx context.Context, otpID string, data *models.OTPData) error {
	data.Tries++
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ttl, err := c.client.TTL(ctx, otpKey(otpID)).Result()
	if err != nil || ttl <= 0 {
		ttl = time.Duration(constant.OTPExpiry) * time.Second
	}
	return c.client.SetEx(ctx, otpKey(otpID), b, ttl).Err()
}

func (c *OTPCache) Delete(ctx context.Context, otpID string) error {
	return c.client.Del(ctx, otpKey(otpID)).Err()
}
