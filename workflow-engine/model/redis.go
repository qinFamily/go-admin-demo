package model

import (
	"context"
	"go-admin-demo/cache"
	"go-admin-demo/tools/config"
	"time"
)

var RedisOpen = config.RedisConfig.RedisType == cache.RedisCluster

// RedisSetVal 将值保存到redis
func RedisSetVal(key, value string, expiration time.Duration) error {
	ctx := context.WithValue(context.Background(), "startKey", time.Now())
	return cache.Redis.Set(ctx, key, value, expiration)
}

// RedisGetVal 从redis获取值
func RedisGetVal(key string) (string, error) {
	ctx := context.WithValue(context.Background(), "startKey", time.Now())
	b, err := cache.Redis.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return string(b), err
}
