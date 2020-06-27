package config

import (
	"time"

	"github.com/spf13/viper"
)

// type redisCache struct {
// 	ring     *redis.Ring
// 	cluster  *redis.ClusterClient
// 	sentinel *redis.Client
// }

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

// const (
// 	redisCluster  = "cluster"
// 	redisRing     = "ring"
// 	redisSentinel = "sentinel"
// )

func InitRedis(r *viper.Viper) *RedisConf {
	// var err error
	return &RedisConf{
		RedisPath:         r.GetString("RedisPath"),
		RedisDatabase:     r.GetInt("RedisDatabase"),
		RedisPassword:     r.GetString("RedisPassword"),
		RedisDialTimeout:  r.GetDuration("RedisDialTimeout"),
		RedisReadTimeout:  r.GetDuration("RedisReadTimeout"),
		RedisWriteTimeout: r.GetDuration("RedisWriteTimeout"),
		RedisPoolSize:     r.GetInt("RedisPoolSize"),
		RedisEnable:       r.GetBool("RedisEnable"),
		RedisType:         r.GetString("RedisType"),
		Expire:            r.GetDuration("Expire"),
		KeyPrefix:         r.GetString("KeyPrefix"),
		ExpireLong:        r.GetDuration("ExpireLong"),
		ExpireShort:       r.GetDuration("ExpireShort"),
		ExpireAssets:      r.GetDuration("ExpireAssets"),

		LruSize:   r.GetInt("LruSize"),
		LruExpire: r.GetDuration("LruExpire"),
	}

	// if RedisConfig.RedisType == redisRing {
	// 	serverList := strings.Split(r.GetString("RedisPath"), ",")
	// 	addrs := make(map[string]string)
	// 	for k, v := range serverList {
	// 		addrs[fmt.Sprintf("redis_%d", k)] = v
	// 	}
	// 	// fmt.Println("addrs", addrs, "RedisPassword", r.GetString("RedisPassword"), "RedisDatabase", r.GetString("RedisDatabase"))
	// 	Redis.ring = redis.NewRing(&redis.RingOptions{
	// 		Addrs:        addrs,
	// 		DialTimeout:  RedisConfig.RedisDialTimeout,
	// 		ReadTimeout:  RedisConfig.RedisReadTimeout,
	// 		WriteTimeout: RedisConfig.RedisWriteTimeout,
	// 		PoolSize:     RedisConfig.RedisPoolSize,
	// 		PoolTimeout:  RedisConfig.Expire,
	// 		Password:     RedisConfig.RedisPassword,
	// 		DB:           RedisConfig.RedisDatabase,
	// 	})

	// 	err = Redis.ring.Ping().Err()
	// } else if RedisConfig.RedisType == redisCluster {
	// 	serverList := strings.Split(r.GetString("RedisPath"), ",")
	// 	Redis.cluster = redis.NewClusterClient(&redis.ClusterOptions{
	// 		Addrs:        serverList,
	// 		DialTimeout:  RedisConfig.RedisDialTimeout,
	// 		ReadTimeout:  RedisConfig.RedisReadTimeout,
	// 		WriteTimeout: RedisConfig.RedisWriteTimeout,
	// 		PoolSize:     RedisConfig.RedisPoolSize,
	// 		PoolTimeout:  RedisConfig.Expire,
	// 		Password:     RedisConfig.RedisPassword,
	// 	})

	// 	err = Redis.cluster.Ping().Err()
	// } else if RedisConfig.RedisType == redisSentinel {
	// 	serverList := strings.Split(r.GetString("RedisPath"), ",")
	// 	Redis.sentinel = redis.NewFailoverClient(&redis.FailoverOptions{
	// 		MasterName:    "mymaster",
	// 		SentinelAddrs: serverList, // A seed list of host:port addresses of sentinel nodes.
	// 		DialTimeout:   RedisConfig.RedisDialTimeout,
	// 		ReadTimeout:   RedisConfig.RedisReadTimeout,
	// 		WriteTimeout:  RedisConfig.RedisWriteTimeout,
	// 		PoolSize:      RedisConfig.RedisPoolSize,
	// 		PoolTimeout:   RedisConfig.Expire,
	// 		Password:      RedisConfig.RedisPassword,
	// 		DB:            RedisConfig.RedisDatabase,
	// 	})

	// 	err = Redis.sentinel.Ping().Err()

	// }
	// return err
}

// var Redis = new(redisCache)
var RedisConfig = new(RedisConf)
