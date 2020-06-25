package cache

import (
	"context"
	"fmt"
	"strings"
	"time"
	"log"
	"go-admin-demo/tools/config"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type redisCache struct {
	ring     *redis.Ring
	cluster  *redis.ClusterClient
	sentinel *redis.Client
}

var Redis redisCache

const (
	redisCluster  = "cluster"
	redisRing     = "ring"
	redisSentinel = "sentinel"
)

func maxlen(data []byte, length int) string {
	if len(data) > length {
		return string(data[:length])
	}
	return string(data)
}

func Init(r *config.RedisConf) error {
	var err error

	if r.RedisType == redisRing {
		serverList := strings.Split(r.RedisPath, ",")
		addrs := make(map[string]string)
		for k, v := range serverList {
			addrs[fmt.Sprintf("redis_%d", k)] = v
		}
		// fmt.Println("addrs", addrs, "RedisPassword", r.RedisPassword, "RedisDatabase", r.RedisDatabase)
		Redis.ring = redis.NewRing(&redis.RingOptions{
			Addrs:        addrs,
			DialTimeout:  r.RedisDialTimeout,
			ReadTimeout:  r.RedisReadTimeout,
			WriteTimeout: r.RedisWriteTimeout,
			PoolSize:     r.RedisPoolSize,
			PoolTimeout:  r.Expire,
			Password:     r.RedisPassword,
			DB:           r.RedisDatabase,
		})

		err = Redis.ring.Ping().Err()
	} else if r.RedisType == redisCluster {
		serverList := strings.Split(r.RedisPath, ",")
		Redis.cluster = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        serverList,
			DialTimeout:  r.RedisDialTimeout,
			ReadTimeout:  r.RedisReadTimeout,
			WriteTimeout: r.RedisWriteTimeout,
			PoolSize:     r.RedisPoolSize,
			PoolTimeout:  r.Expire,
			Password:     r.RedisPassword,
		})

		err = Redis.cluster.Ping().Err()
	} else if r.RedisType == redisSentinel {
		serverList := strings.Split(r.RedisPath, ",")
		Redis.sentinel = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "mymaster",
			SentinelAddrs: serverList, // A seed list of host:port addresses of sentinel nodes.
			DialTimeout:   r.RedisDialTimeout,
			ReadTimeout:   r.RedisReadTimeout,
			WriteTimeout:  r.RedisWriteTimeout,
			PoolSize:      r.RedisPoolSize,
			PoolTimeout:   r.Expire,
			Password:      r.RedisPassword,
			DB:            r.RedisDatabase,
		})

		err = Redis.sentinel.Ping().Err()

	}

	return err
}

func (c *redisCache) Get(ctx context.Context, key string) ([]byte, error) {
	var data []byte
	var err error
	now := time.Now()
	if c.ring != nil {
		data, err = c.ring.Get(key).Bytes()
	} else if c.cluster != nil {
		data, err = c.cluster.Get(key).Bytes()
	} else if c.sentinel != nil {
		data, err = c.sentinel.Get(key).Bytes()
	}
	elasped := time.Since(now)
	// start := ctx.Value("startKey").(time.Time)

	if err != nil {
		if err != redis.Nil {
			// metrics\.RedisCounter.WithLabelValues("fail").Inc()
			log.Println("redisCache Get")
			// metrics\.SystemError.WithLabelValues(metrics.RedisGetError).Inc()
		} else {
			// metrics\.RedisCounter.WithLabelValues("miss").Inc()
			log.Println("redisCache Get")
		}
	} else {
		// metrics\.RedisCounter.WithLabelValues("hit").Inc()
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache Get spend too long")
		} else {
			log.Println("redisCache Get")
		}
	}

	return data, err
}

func (c *redisCache) Set(ctx context.Context, key string, val string, expiration time.Duration) error {
	var err error
	now := time.Now()

	if c.ring != nil {
		err = c.ring.Set(key, val, expiration).Err()
	} else if c.cluster != nil {
		err = c.cluster.Set(key, val, expiration).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.Set(key, val, expiration).Err()
	}


	if err != nil && err != redis.Nil {
		log.Println("redisCache Set")
		// metrics\.SystemError.WithLabelValues(metrics.RedisSetError).Inc()
	} else {

		// 大于0.5秒 记录错误日志
		if time.Since(now) >= time.Millisecond*200 {
			log.Println("redisCache Set too long")
		} else {
			log.Println("redisCache Set")
		}
	}

	return err
}

func (c *redisCache) Delete(ctx context.Context, key string) error {
	var err error
	now := time.Now()
	if c.ring != nil {
		err = c.ring.Del(key).Err()
	} else if c.cluster != nil {
		err = c.cluster.Del(key).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.Del(key).Err()
	}

	start := ctx.Value("startKey").(time.Time)

	elasped := time.Since(now)
	flowElasped := time.Since(start)
	fields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "delete"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache Del", fields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache Del too long", fields)
		} else {
			log.Println("redisCache Del", fields)
		}
	}

	return err
}

func (c *redisCache) LPUSH(ctx context.Context, key string, val string) error {
	var err error
	now := time.Now()

	if c.ring != nil {
		err = c.ring.LPush(key, val).Err()
	} else if c.cluster != nil {
		err = c.cluster.LPush(key, val).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.LPush(key, val).Err()
	}
	elasped := time.Since(now)
	if err != nil && err != redis.Nil {
		log.Println("redisCache LPUSH")
		// metrics\.SystemError.WithLabelValues(metrics.RedisSetError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache LPUSH spent too long")
		} else {
			log.Println("redisCache LPUSH finished")
		}
	}
	return err
}

func (c *redisCache) INCR(ctx context.Context, key string) error {
	var err error
	now := time.Now()
	if c.ring != nil {
		err = c.ring.Incr(key).Err()
	} else if c.cluster != nil {
		err = c.cluster.Incr(key).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.Incr(key).Err()
	}
	start := ctx.Value("startKey").(time.Time)

	elasped := time.Since(now)
	flowElasped := time.Since(start)
	fields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "Incr"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache Incr", fields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisIncrError).Inc()
	} else {
		log.Println("redisCache Incr", fields)
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache INCR spend too long", fields)
		} else {
			log.Println("redisCache INCR finished", fields)
		}
	}

	return err
}

func (c *redisCache) RPop(ctx context.Context, key string) (string, error) {
	var err error
	var val string
	now := time.Now()

	if c.ring != nil {
		val, err = c.ring.RPop(key).Result()
	} else if c.cluster != nil {
		val, err = c.cluster.RPop(key).Result()
	} else if c.sentinel != nil {
		err = c.sentinel.RPop(key).Err()
	}
	elasped := time.Since(now)
	if err != nil && err != redis.Nil {
		log.Println("redisCache RPop")
		// metrics\.SystemError.WithLabelValues(metrics.RedisSetError).Inc()
	} else {
		log.Println("redisCache RPop")
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache RPop spent too long")
		} else {
			log.Println("redisCache RPop finished")
		}
	}
	return val, err
}

func (c *redisCache) MGet(ctx context.Context, key string) (data []interface{}, err error) {

	now := time.Now()
	if c.ring != nil {
		data, err = c.ring.MGet(key).Result()
	} else if c.cluster != nil {
		data, err = c.cluster.MGet(key).Result()
	} else if c.sentinel != nil {
		data, err = c.sentinel.MGet(key).Result()
	}
	start := ctx.Value("startKey").(time.Time)

	elasped := time.Since(now)
	flowElasped := time.Since(start)
	fields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "MGet"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache MGet", fields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache MGet too long", fields)
		} else {
			log.Println("redisCache MGet", fields)
		}
	}

	return
}

func (c *redisCache) MSet(ctx context.Context, key string) error {
	var err error
	now := time.Now()
	start := ctx.Value("startKey").(time.Time)
	if c.ring != nil {
		err = c.ring.MSet(key).Err()
	} else if c.cluster != nil {
		err = c.cluster.MSet(key).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.MSet(key).Err()
	}
	elasped := time.Since(now)
	flowElasped := time.Since(start)
	fields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "MSet"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache MSet", fields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache MSet too long", fields)
		} else {
			log.Println("redisCache MSet", fields)
		}
	}

	return err
}

func (c *redisCache) HMGet(ctx context.Context, key string, fields ...string) (data []interface{}, err error) {

	now := time.Now()
	if c.ring != nil {
		data, err = c.ring.HMGet(key, fields...).Result()
	} else if c.cluster != nil {
		data, err = c.cluster.HMGet(key, fields...).Result()
	} else if c.sentinel != nil {
		data, err = c.sentinel.HMGet(key, fields...).Result()
	}
	start := ctx.Value("startKey").(time.Time)

	elasped := time.Since(now)
	flowElasped := time.Since(start)
	zapfields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Any("fields", fields),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "HMGet"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache HMGet", zapfields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache HMGet too long", zapfields)
		} else {
			log.Println("redisCache HMGet", zapfields)
		}
	}

	return
}

func (c *redisCache) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	var err error
	now := time.Now()
	start := ctx.Value("startKey").(time.Time)
	if c.ring != nil {
		err = c.ring.HMSet(key, fields).Err()
	} else if c.cluster != nil {
		err = c.cluster.HMSet(key, fields).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.HMSet(key, fields).Err()
	}
	elasped := time.Since(now)
	flowElasped := time.Since(start)
	zapfields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Any("fields", fields),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "HMSet"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache HMSet", zapfields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache HMSet too long", zapfields)
		} else {
			log.Println("redisCache HMSet", zapfields)
		}
	}

	return err
}

func (c *redisCache) HGETALL(ctx context.Context, key string) (data map[string]string, err error) {

	now := time.Now()
	if c.ring != nil {
		data, err = c.ring.HGetAll(key).Result()
	} else if c.cluster != nil {
		data, err = c.cluster.HGetAll(key).Result()
	} else if c.sentinel != nil {
		data, err = c.sentinel.HGetAll(key).Result()
	}
	start := ctx.Value("startKey").(time.Time)

	elasped := time.Since(now)
	flowElasped := time.Since(start)
	fields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "HGetAll"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache HGetAll", fields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache HGetAll too long", fields)
		} else {
			log.Println("redisCache HGetAll", fields)
		}
	}

	return
}

func (c *redisCache) HDel(ctx context.Context, key string, fields ...string) error {
	var err error
	now := time.Now()
	start := ctx.Value("startKey").(time.Time)
	if c.ring != nil {
		err = c.ring.HDel(key, fields...).Err()
	} else if c.cluster != nil {
		err = c.cluster.HDel(key, fields...).Err()
	} else if c.sentinel != nil {
		err = c.sentinel.HDel(key, fields...).Err()
	}
	elasped := time.Since(now)
	flowElasped := time.Since(start)
	zapfields := []zapcore.Field{
		zap.Any("cacheResponseTime", elasped.String()),
		zap.Any("flowTime", flowElasped.String()),
		zap.String("query", key),
		zap.Any("fields", fields),
		zap.Error(err),
		zap.String("type", "redis"),
		zap.String("action", "HDel"),
		zap.String("direction", "out")}

	if err != nil && err != redis.Nil {
		log.Println("redisCache HDel", zapfields)
		// metrics\.SystemError.WithLabelValues(metrics.RedisDelError).Inc()
	} else {
		// 大于0.2秒 记录错误日志
		if elasped >= time.Millisecond*200 {
			log.Println("redisCache HDel too long", zapfields)
		} else {
			log.Println("redisCache HDel", zapfields)
		}
	}

	return err
}
