package cache

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	// "database/sql"

	"github.com/golang/groupcache/singleflight"
	lru "github.com/hnlq715/golang-lru"
)

type lruCache struct {
	maxSize int
	expire  time.Duration

	g   singleflight.Group
	arc *lru.ARCCache
}

var Nil = errors.New("lruCache: nil")

const (
	LruSize   = 200
	LruExpire = time.Hour * 168
)

type GetterFunc func() (interface{}, error)

var cache *lruCache
var cacheExpire map[string]*lruCache
var cacheMux sync.RWMutex

func LRU() *lruCache {
	cacheMux.RLock()
	defer cacheMux.RUnlock()
	if cache != nil {
		return cache
	}
	cache = NewARC(LruSize, LruExpire)
	return cache
}

func NewARC(maxSize int, expire time.Duration) *lruCache {
	arc, _ := lru.NewARCWithExpire(maxSize, expire)

	return &lruCache{
		maxSize: maxSize,
		expire:  expire,
		arc:     arc,
	}
}

func (l *lruCache) Get(key string) (interface{}, error) {
	data, ok := l.arc.Get(key)
	if !ok {
		return nil, Nil
	}
	return data, nil
}

func (l *lruCache) Del(key string) {
	l.arc.Remove(key)
}

func (l *lruCache) GetWithLoader(key string, load GetterFunc) (interface{}, error) {
	data, ok := l.arc.Get(key)
	if !ok && load != nil {
		// metrics.LruCounter.WithLabelValues("miss").Inc()
		return l.g.Do(key, func() (interface{}, error) {
			now := time.Now()
			data, err := load()
			elasped := time.Since(now)
			if elasped > 10*time.Second {
				log.Println(fmt.Sprintf(`{"ts":"%s","msg":"lru get spend too long","query":"%s", "lruResponseTime":%+v}`, time.Now().Local().Format(time.RFC1123), key, elasped.String()))
			}
			// metrics.LruLoadLatency.Observe(elasped.Seconds())
			if err != nil {
				// if err != sql.ErrNoRows {
				// metrics.SystemError.WithLabelValues(metrics.LruLoadError).Inc()
				// metrics.LruCounter.WithLabelValues("fail").Inc()
				// }
				// log.Println("=============================== lru error", err)
				return nil, err
			}
			// metrics.LruCounter.WithLabelValues("load").Inc()
			// log.Println("===============================", fmt.Sprintf("key=>%s,  data=>%+v", key, data))
			err = l.Set(key, data)
			return data, err
		})
	}

	// metrics.LruCounter.WithLabelValues("hit").Inc()
	// log.Println("lru get succeed", zap.Any("query", key), zap.String("type", "lru"), zap.Any("data", data))
	return data, nil
}

func (l *lruCache) Set(key string, data interface{}) error {
	l.arc.Add(key, data)
	return nil
}

func LRUWithExpire(expire time.Duration) *lruCache {
	cacheMux.RLock()
	defer cacheMux.RUnlock()
	expireStr := fmt.Sprintf("%d", int64(expire))
	if cacheExpire == nil {
		cacheExpire = make(map[string]*lruCache)
	}
	if cacheExpire[expireStr] != nil {
		return cacheExpire[expireStr]
	}
	cacheExpire[expireStr] = NewARC(LruSize, expire)
	return cacheExpire[expireStr]
}
