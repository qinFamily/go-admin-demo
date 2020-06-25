package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type redisCache struct {
	ring     *redis.Ring
	cluster  *redis.ClusterClient
	sentinel *redis.Client
}

type RedisConf struct {
	RedisPath         string
	RedisDatabase     int
	RedisPassword     string
	RedisDialTimeout  time.Duration
	RedisReadTimeout  time.Duration
	RedisWriteTimeout time.Duration
	RedisPoolSize     int
	RedisEnable       bool
	RedisType         string
	Expire            time.Duration
	KeyPrefix         string
	ExpireLong        time.Duration
	ExpireShort       time.Duration
	ExpireAssets      time.Duration

	LruSize   int
	LruExpire time.Duration
}

const (
	redisCluster  = "cluster"
	redisRing     = "ring"
	redisSentinel = "sentinel"
)

func InitRedis(r *viper.Viper) error {
	var err error

	if r.GetString("Type") == redisRing {
		serverList := strings.Split(r.GetString("RedisPath"), ",")
		addrs := make(map[string]string)
		for k, v := range serverList {
			addrs[fmt.Sprintf("redis_%d", k)] = v
		}
		// fmt.Println("addrs", addrs, "RedisPassword", r.GetString("RedisPassword"), "RedisDatabase", r.GetString("RedisDatabase"))
		Redis.ring = redis.NewRing(&redis.RingOptions{
			Addrs:        addrs,
			DialTimeout:  r.GetDuration("RedisDialTimeout"),
			ReadTimeout:  r.GetDuration("RedisReadTimeout"),
			WriteTimeout: r.GetDuration("RedisWriteTimeout"),
			PoolSize:     r.GetInt("RedisPoolSize"),
			PoolTimeout:  r.GetDuration("Expire"),
			Password:     r.GetString("RedisPassword"),
			DB:           r.GetInt("RedisDatabase"),
		})

		err = Redis.ring.Ping().Err()
	} else if r.GetString("RedisType") == redisCluster {
		serverList := strings.Split(r.GetString("RedisPath"), ",")
		Redis.cluster = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        serverList,
			DialTimeout:  r.GetDuration("RedisDialTimeout"),
			ReadTimeout:  r.GetDuration("RedisReadTimeout"),
			WriteTimeout: r.GetDuration("RedisWriteTimeout"),
			PoolSize:     r.GetInt("RedisPoolSize"),
			PoolTimeout:  r.GetDuration("Expire"),
			Password:     r.GetString("RedisPassword"),
		})

		err = Redis.cluster.Ping().Err()
	} else if r.GetString("RedisType") == redisSentinel {
		serverList := strings.Split(r.GetString("RedisPath"), ",")
		Redis.sentinel = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "mymaster",
			SentinelAddrs: serverList, // A seed list of host:port addresses of sentinel nodes.
			DialTimeout:  r.GetDuration("RedisDialTimeout"),
			ReadTimeout:  r.GetDuration("RedisReadTimeout"),
			WriteTimeout: r.GetDuration("RedisWriteTimeout"),
			PoolSize:     r.GetInt("RedisPoolSize"),
			PoolTimeout:  r.GetDuration("Expire"),
			Password:     r.GetString("RedisPassword"),
			DB:           r.GetInt("RedisDatabase"),
		})

		err = Redis.sentinel.Ping().Err()

	}
	return err
}

var Redis = new(redisCache)
